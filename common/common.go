package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/qiniu/log.v1"
)

const (
	VariableRetentionPolicyNameFormat = "%s_rp"

	RetentionPolicyForever = "default"
)

type QueryRet struct {
	Results []Result `json:"results,omitempty"`
	Err     error    `json:"error,omitempty"`
}

func (e QueryRet) String() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		log.Error(err)
		return ""
	}

	return string(bytes)
}

type Result struct {
	Series []Serie `json:"series,omitempty"`
	Err    string  `json:"error,omitempty"`
}

func (e Result) String() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		log.Error(err)
		return ""
	}

	return string(bytes)
}

type Serie struct {
	Name    string            `json:"name"`
	Columns []string          `json:"columns"`
	Tags    map[string]string `json:"tags"`
	Values  [][]interface{}   `json:"values"`
}

func (e Serie) String() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		log.Error(err)
		return ""
	}

	return string(bytes)
}

func WritePoints(
	host string, port int, db string, rp string, points string) error {

	addr := host + ":" + strconv.Itoa(port)
	url := "http://" + addr + "/write?db=" + db + "&rp=" + string(rp)
	log.Info("Write points query:", url, ":", points)
	body := strings.NewReader(points)
	resp, err := http.Post(url, "text/plain", body)
	defer CloseResponse(resp)
	
	log.Info(resp.Status)

	return err
}

func QueryInfluxdb(host string, port int, db string, sql string) (ret QueryRet, err error) {
	addr := host + ":" + strconv.Itoa(port)
	queryURL := "http://" + addr + "/query?db=" + db + "&q=" + url.QueryEscape(sql)
	resp, err := http.Get(queryURL)
	defer CloseResponse(resp)
	
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("Request [%v] Error: %v", queryURL, resp)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &ret)
	if err != nil {
		return
	}

	return
}

func CreatDatabase(host string, port int, dbName string) error {
	addr := host + ":" + strconv.Itoa(port)
	url := "http://" + addr + "/query?q=CREATE+DATABASE+" + dbName
	log.Info(url)
	resp, err := http.Get(url)
	defer CloseResponse(resp)

	if err != nil {
		return fmt.Errorf("%v Create database fail:%v\n", url, err)
	}
	if resp != nil {
		_, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf(
				"%v Create database fail,respcode: %v, error: %v",
				url,
				resp.StatusCode,
				err,
			)
		}
	}
	
	return nil
}

// CREATE RETENTION POLICY variable_rp ON mydb DURATION 0s REPLICATION 1
func CreatRetentionPolicy(
	host string,
	port int,
	dbName string,
	rpName string,
	rp string) (ret QueryRet, err error) {

	var sql string
	if rp != "" || rp == RetentionPolicyForever {
		sql = fmt.Sprintf(
			"CREATE RETENTION POLICY %s ON %s DURATION 0s REPLICATION 1",
			rpName,
			dbName,
		)
	} else {
		sql = fmt.Sprintf(
			"CREATE RETENTION POLICY %s ON %s DURATION %s REPLICATION 1",
			rpName,
			dbName,
			rp,
		)
	}

	return QueryInfluxdb(host, port, dbName, sql)
}

func AlterRetentionPolicyToMinDuration(
	host string,
	port int,
	dbName string,
	rpName string) (ret QueryRet, err error) {

	sql := fmt.Sprintf(
		"ALTER RETENTION POLICY %s on %s duration 1h",
		rpName,
		dbName,
	)

	return QueryInfluxdb(host, port, dbName, sql)
}

func CloseResponse(resp *http.Response) {
	if resp == nil {
		return
	}
	err := resp.Body.Close()
	if err != nil {
		log.Error(err)
	}
}

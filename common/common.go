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

type QueryRet struct {
	Results []Result `json:"results,omitempty"`
	Err     error      `json:"error,omitempty"`
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
	Err    string    `json:"error,omitempty"`
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
	host string, port int, db string, points string) error {

	addr := host + ":" + strconv.Itoa(port)
	url := "http://" + addr + "/write?db=" + db
	log.Info("Write points query:", points)
	body := strings.NewReader(points)
	_, err := http.Post(url, "text/plain", body)

	return err
}

func QueryInfluxdb(host string, port int, db string, sql string) (ret QueryRet, err error) {
	
	addr := host + ":" + strconv.Itoa(port)
	queryURL := "http://" + addr + "/query?db=" + db + "&q=" + url.QueryEscape(sql)
	resp, err := http.Get(queryURL)
	if resp != nil {
		defer resp.Body.Close()
	}
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

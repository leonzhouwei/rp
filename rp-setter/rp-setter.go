package main

import (
	"fmt"
	"time"

	"tmp/common"
	"tmp/conf"

	"github.com/qiniu/log.v1"
)

const (
	sleepDuration = time.Second
)

func init() {
	log.SetOutputLevel(log.Ldebug)
}

func main() {
	host := conf.Host
	port := conf.Port
	db := conf.Db
	min := conf.MinSeriesNum
	max := conf.MaxSeriesNum

	// create the database if not exits
	err := common.CreatDatabase(host, port, db)
	if err != nil {
		log.Fatal(err)
	}

	// write the RP-variable series
	for i := min; i <= max; i++ {
		seriesName := fmt.Sprintf("cpu_%v", i)
		// create the RP
		rpName := fmt.Sprintf(
			common.VariableRetentionPolicyNameFormat,
			seriesName,
		)

		// alter rp
		msg := fmt.Sprintf(
			"alter %s of %s.%s to minimum duration", 
			rpName,
			db,
			seriesName,
		)
		result, err := common.AlterRetentionPolicyToMinDuration(
			host,
			port,
			db,
			rpName,
		)
		if err != nil {
			log.Fatal(msg, "failed:", err)
		} else {
			log.Info(msg, "ok:", result)
		}
		time.Sleep(sleepDuration)
	}
}

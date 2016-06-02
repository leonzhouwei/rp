package main

import (
	"fmt"
	"time"

	"tmp/common"
	"tmp/conf"

	"github.com/qiniu/log.v1"
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

	// write the RP-variable series
	for i := min; i <= max; i++ {
		seriesName := fmt.Sprintf("cpu_%v", i)
		// create the RP
		rpName := fmt.Sprintf(
			common.VariableRetentionPolicyNameFormat,
			seriesName,
		)

		fmt.Println(
			"i will count 10 seconds before the benchmark after changing the duration of the RP",
			rpName,
			"of",
			seriesName,
		)
		for i := 9; i >= 0; i-- {
			fmt.Println(i)
			time.Sleep(time.Second * 1)
		}

		// alter rp
		msg := fmt.Sprintf(
			"alter %s of %s.%s to minimum duration",
			rpName,
			db,
			seriesName,
		)
		result, err := common.AlterRetentionPolicyDurationToOneHour(
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
	}
}

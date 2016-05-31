package main

import (
	"fmt"
	"time"

	"tmp/common"
	"tmp/conf"

	"github.com/qiniu/log.v1"
)

const (
	points = `cpu_%v,host=host_%v id=%v,created_at="%v"`
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
		result, err := common.CreatRetentionPolicy(
			host,
			port,
			db,
			rpName,
			common.RetentionPolicyForever,
		)
		if err != nil {
			log.Fatal(err)
		}
		log.Info(result)

		// write the points
		for j := 1; j <= 1000; j++ {
			pointsString := fmt.Sprintf(
				points,
				i,
				i,
				j,
				time.Now().Format(time.RFC3339Nano),
			)
			err := common.WritePoints(
				host,
				port,
				db,
				rpName,
				pointsString,
			)
			if err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
}

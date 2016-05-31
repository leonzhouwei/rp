package main

import (
	"time"
	"fmt"

	"tmp/common"
	"tmp/conf"
	
	"github.com/qiniu/log.v1"
)

const (
	points = `cpu_%v,host=host_%v value=%v,time_human="%v"`
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

	// write
	for i := min; i <= max; i++ {
		pointsString := fmt.Sprintf(points, i, i, i, time.Now().Format(time.RFC3339Nano))
		err := common.WritePoints(host, port, db, common.SystemRp1Hour, pointsString)
		if err != nil {
			log.Fatal(err)
		}
	}
}

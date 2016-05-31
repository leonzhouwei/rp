package main

import (
	"fmt"
	"time"

	"tmp/common"
	"tmp/conf"

	"github.com/qiniu/log.v1"
)

const (
	points = `cpu_%v,host=host_%v value=%v,created_at="%v"`
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
		err := common.WritePoints(host, port, db, common.RetentionPolicyForever, pointsString)
		if err != nil {
			log.Fatal(err)
		}
	}
}

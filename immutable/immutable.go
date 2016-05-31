package main

import (
	"fmt"
	"time"

	"tmp/common"
	"tmp/conf"

	"github.com/qiniu/log.v1"
)

const (
	points = `immutable,host=host_%v id=%v,created_at="%v"`
)

func init() {
	log.SetOutputLevel(log.Ldebug)
}

func main() {
	host := conf.Host
	port := conf.Port
	db := conf.Db

	// write
	for i := 1; i <= 10000; i++ {
		pointsString := fmt.Sprintf(
			points,
			i,
			i,
			time.Now().Format(time.RFC3339Nano),
		)
		err := common.WritePoints(
			host,
			port,
			db,
			common.RetentionPolicyForever,
			pointsString,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}

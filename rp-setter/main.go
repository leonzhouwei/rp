package main

import (
	"tmp/common"
	"tmp/conf"

	"github.com/qiniu/log.v1"
)

const (
	sql = "ALTER RETENTION POLICY " + common.RetentionPolicyVariable + " on mydb duration 1h"
)

func init() {
	log.SetOutputLevel(log.Ldebug)
}

func main() {
	host := conf.Host
	port := conf.Port
	db := conf.Db

	// alter rp
	result, err := common.QueryInfluxdb(host, port, db, sql)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(result)
}

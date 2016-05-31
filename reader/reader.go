package main

import (
	"tmp/common"
	"tmp/conf"
	
	"github.com/qiniu/log.v1"
)

const (
	sql = "select id from cpu_1"
)

func main() {
	host := conf.Host
	port := conf.Port
	db := conf.Db

	// read
	result, err := common.QueryInfluxdb(host, port, db, sql)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(result)
}

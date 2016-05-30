package main

import (
	"github.com/qiniu/log.v1"
	"tmp/common"
)

const (
	sql = "select value from cpu_1"
)

func main() {
	host := "127.0.0.1"
	port := 8086
	db := "mydb"

	// read
	result, err := common.QueryInfluxdb(host, port, db, sql)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(result)
}

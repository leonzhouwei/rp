package main

import (
	"github.com/qiniu/log.v1"
	"tmp/common"
)

const (
	sql    = "select value from cpu"
	points = "cpu,host=s1 value=10 1"
)

func main() {
	host := "127.0.0.1"
	port := 8086
	db := "mydb"

	// write
	err := common.WritePoints(host, port, db, points)
	if err != nil {
		log.Fatal(err)
	}

	// read
	result, err := common.QueryInfluxdb(host, port, db, sql)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(result)
}

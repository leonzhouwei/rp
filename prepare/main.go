package main

import (
	"fmt"

	"github.com/qiniu/log.v1"
	"tmp/common"
)

const (
	points = "cpu_%v,host=host_%v value=%v %v"
)

func main() {
	host := "127.0.0.1"
	port := 8086
	db := "mydb"

	// write
	for i := common.MinSeriesNum; i <= common.MaxSeriesNum; i++ {
		pointsString := fmt.Sprintf(points, i, i, i, i)
		err := common.WritePoints(host, port, db, pointsString)
		if err != nil {
			log.Fatal(err)
		}
	}
}

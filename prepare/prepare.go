package main

import (
	"fmt"

	"tmp/common"
	"tmp/conf"
	
	"github.com/qiniu/log.v1"
)

const (
	points = "cpu_%v,host=host_%v value=%v %v"
)

func main() {
	host := conf.Host
	port := conf.Port
	db := conf.Db
	min := conf.MinSeriesNum
	max := conf.MaxSeriesNum

	// write
	for i := min; i <= max; i++ {
		pointsString := fmt.Sprintf(points, i, i, i, i)
		err := common.WritePoints(host, port, db, pointsString)
		if err != nil {
			log.Fatal(err)
		}
	}
}

package conf

const (
	MinSeriesNum = 1

	// 我的 15 寸 rmbp 上只能创建 243 个 series，且
	// 每个 series 都有一个独立的 RP
	// 这里的 241 == 243 - 1 - 1，是因为
	// 还有一个名为 immutable 的 series 要写到 influxdb 里面
	// 另外还有一个名为 req 的 series 预留给孟总给我的 export.bug (2.1GB) 
	MaxSeriesNum = 1
	Host         = "127.0.0.1"
	Port         = 8086
	Db           = "mydb"
)

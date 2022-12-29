package collector

import (
	pb "cmdb-agent/proto"

	"github.com/shirou/gopsutil/load"
)

func collectLoadInfo() pb.LoadAvg {
	loadAvg, _ := load.Avg()
	return pb.LoadAvg{
		Load_1Min:  float32(loadAvg.Load1),
		Load_5Min:  float32(loadAvg.Load5),
		Load_15Min: float32(loadAvg.Load15),
	}
}

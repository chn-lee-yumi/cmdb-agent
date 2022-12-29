package collector

import (
	pb "cmdb-agent/proto"

	"github.com/shirou/gopsutil/cpu"
)

func collectCpuInfo() pb.CpuInfo {
	//cpu基本信息
	cpuInfos, err := cpu.Info()
	if err != nil {
		panic(err)
	}
	cpuName := cpuInfos[0].ModelName
	//cpu数量;true逻辑核心数量，false物理核心数量
	cpuLogicalCount, err := cpu.Counts(true)
	if err != nil {
		panic(err)
	}
	cpuPhysicalCount, err := cpu.Counts(false)
	if err != nil {
		panic(err)
	}
	//cpu利用率;true为每个cpu，false为总的cpu
	// cpuUsage, err := cpu.Percent(time.Second, true)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(cpuUsage)
	//cpu有关时间信息;true为每个cpu，false为总的cpu
	// cpuTime, err := cpu.Times(true)
	// cpu := Cpu{
	// 	Info:          cpuInfos,
	// 	LogicalCount:  cpuLogicalCount,
	// 	PhysicalCount: cpuPhysicalCount,
	// 	Usage:         cpuUsage,
	// 	Time:          cpuTime,
	// }
	return pb.CpuInfo{
		Name:          cpuName,
		LogicalCount:  uint32(cpuLogicalCount),
		PhysicalCount: uint32(cpuPhysicalCount),
	}
}

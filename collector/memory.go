package collector

import (
	pb "cmdb-agent/proto"

	"github.com/shirou/gopsutil/mem"
)

func collectMemoryInfo() pb.MemoryInfo {
	physicalMemory, _ := mem.VirtualMemory()
	swapMemory, _ := mem.SwapMemory()
	return pb.MemoryInfo{
		PhysicalMemoryTotal:       physicalMemory.Total,
		PhysicalMemoryUsedPercent: float32(physicalMemory.UsedPercent),
		SwapMemoryTotal:           swapMemory.Total,
		SwapMemoryUsedPercent:     float32(swapMemory.UsedPercent),
	}
}

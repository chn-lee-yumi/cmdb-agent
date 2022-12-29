package collector

import pb "cmdb-agent/proto"

func Collect() pb.Data {
	// 采集数据 https://blog.csdn.net/shn111/article/details/122388723
	deviceSystemInfo := collectDeviceSystemInfo()
	cpuInfo := collectCpuInfo()
	systemInfo := collectSystemInfo()
	memoryInfo := collectMemoryInfo()
	loadAvg := collectLoadInfo()
	// collectDiskInfo()
	networkInfo := collectNetworkInfo()
	// 组装并返回
	return pb.Data{
		DeviceSystemInfo: &deviceSystemInfo,
		SystemInfo:       &systemInfo,
		CpuInfo:          &cpuInfo,
		MemoryInfo:       &memoryInfo,
		LoadAvg:          &loadAvg,
		NetworkInfo:      &networkInfo,
	}
}

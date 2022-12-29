package collector

import (
	pb "cmdb-agent/proto"

	"github.com/shirou/gopsutil/host"
)

func collectSystemInfo() pb.SystemInfo {
	info, err := host.Info()
	if err != nil {
		return pb.SystemInfo{}
	}
	return pb.SystemInfo{
		Hostname:           info.Hostname,
		Uptime:             info.Uptime,
		Os:                 info.OS,
		Platform:           info.Platform,
		PlatformVersion:    info.PlatformVersion,
		KernelVersion:      info.KernelVersion,
		KernelArch:         info.KernelArch,
		VirtualizationRole: info.VirtualizationRole,
	}
}

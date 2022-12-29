package collector

import (
	"log"

	"github.com/shirou/gopsutil/disk"
)

func collectDiskInfo() {
	//TODO: 获取磁盘分区 https://pkg.go.dev/github.com/shirou/gopsutil@v2.21.11+incompatible/disk#Partitions
	DiskParti, _ := disk.Partitions(false)
	log.Println(DiskParti)
}

package collector

import (
	pb "cmdb-agent/proto"
	"runtime"
)

func collectDeviceSystemInfo() pb.DeviceSystemInfo {
	var productName, manufacturer, serialNumber, version string
	if runtime.GOOS == "linux" {
		productName = execShell("dmidecode -s system-product-name")
		manufacturer = execShell("dmidecode -s system-manufacturer")
		serialNumber = execShell("dmidecode -s system-serial-number")
		version = execShell("dmidecode -s system-version")
	}
	return pb.DeviceSystemInfo{
		ProductName:  productName,
		Manufacturer: manufacturer,
		SerialNumber: serialNumber,
		Version:      version,
	}
}

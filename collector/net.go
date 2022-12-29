package collector

import (
	"fmt"
	"log"
	"net"
	"runtime"

	pb "cmdb-agent/proto"
)

func getMainIp() string {
	var mainIp string //默认路由对应网卡的第一个IP
	if runtime.GOOS == "linux" {
		// 获取默认路由
		output := execShell("ip route show default")
		// 解析路由输出，提取网卡名称
		var gateway, dev, interfaceName string
		_, err := fmt.Sscanf(output, "default via %s %s %s", &gateway, &dev, &interfaceName)
		if err != nil {
			log.Println("Error parsing route output:", err)
			return ""
		}
		// 获取网卡信息
		interfaceAddr, err := net.InterfaceByName(interfaceName)
		if err != nil {
			log.Println("Error getting interface:", err)
			return ""
		}
		// 获取网卡地址
		addresses, err := interfaceAddr.Addrs()
		if err != nil {
			log.Println("Error getting addresses:", err)
			return ""
		}
		// 取第一个网卡地址
		for _, address := range addresses {
			mainIp = address.String()
			break
		}
	}
	return mainIp
}

func getInterfaces() []*pb.Interface {
	interfaces := []*pb.Interface{}
	// 获取所有网卡
	netInterfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalln(err)
		return interfaces
	}
	// 遍历网卡
	for _, i := range netInterfaces {
		ips := []string{}
		// 获取IP列表
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			ips = append(ips, addr.String())
		}
		interfaces = append(interfaces,
			&pb.Interface{
				Name:   i.Name,
				MTU:    uint32(i.MTU),
				Hwaddr: i.HardwareAddr.String(),
				IPs:    ips,
			})
	}
	return interfaces
}

func collectNetworkInfo() pb.NetworkInfo {
	return pb.NetworkInfo{
		MainIp:     getMainIp(),
		Interfaces: getInterfaces(),
	}
}

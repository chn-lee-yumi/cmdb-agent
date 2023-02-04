package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	"cmdb-agent/collector"
	pb "cmdb-agent/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
)

var grpcServers = [...]string{"cmdb.gcc.ac.cn:8083", "cmdb.gcc.ac.cn:8084"} // grpc服务器列表，默认请求第一个，第一个连不上后会请求下一个

func main() {
	// 自动更新
	updateUrl := fmt.Sprintf("http://download.gcc.ac.cn/cmdb-agent-%s-%s", runtime.GOOS, runtime.GOARCH)
	log.Println("updateUrl: ", updateUrl)
	overseer.Run(overseer.Config{
		Program: prog,
		Fetcher: &fetcher.HTTP{
			URL:      updateUrl,
			Interval: 1 * time.Minute,
		},
	})
}

// 真正的main
func prog(state overseer.State) {
	log.Printf("app (%s) running...", state.ID)

	// 创建grpc连接
	clientList := make([]pb.ReceiverClient, 0, len(grpcServers))
	for i := 0; i < len(grpcServers); i++ {
		conn, err := grpc.Dial(grpcServers[i], grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Connect error: %v", err)
		}
		defer conn.Close()
		c := pb.NewReceiverClient(conn)
		clientList = append(clientList, c)
	}

	nextExecTime := time.Now() //下次执行时间
	for {
		// 计算下次执行的时间
		nextExecTime = nextExecTime.Add(time.Minute)
		// 采集数据
		log.Println("collecting")
		p := collector.Collect()
		// 上传
		log.Println("uploading")
		isSuccess := false
		for i := 0; i < len(grpcServers); i++ {
			ack, err := clientList[i].Post(context.Background(), &p)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("return code:", ack.Code)
				isSuccess = true
				break
			}
		}
		if isSuccess {
			log.Println("upload success")
		} else {
			log.Println("upload failed, all servers down")
		}
		// 等待下次采集
		time.Sleep(nextExecTime.Sub(time.Now()))
	}
}

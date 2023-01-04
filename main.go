package main

import (
	"context"
	"log"
	"time"

	"cmdb-agent/collector"
	pb "cmdb-agent/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
)

func main() {
	// 自动更新
	overseer.Run(overseer.Config{
		Program: prog,
		Fetcher: &fetcher.HTTP{
			URL:      "http://download.gcc.ac.cn/cmdb-agent",
			Interval: 1 * time.Minute,
		},
	})
}

// 真正的main
func prog(state overseer.State) {
	log.Printf("app (%s) running...", state.ID)

	// 连接grpc
	conn, err := grpc.Dial("cmdb.gcc.ac.cn:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connect error: %v", err)
	}
	defer conn.Close()
	c := pb.NewReceiverClient(conn)
	nextExecTime := time.Now()

	for {
		// 计算下次执行的时间
		nextExecTime = nextExecTime.Add(time.Minute)
		// 采集数据
		log.Println("collecting")
		p := collector.Collect()
		// 上传
		log.Println("uploading")
		ack, err := c.Post(context.Background(), &p)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("return code:", ack.Code)
		}
		// 等待下次采集
		time.Sleep(nextExecTime.Sub(time.Now()))
	}
}

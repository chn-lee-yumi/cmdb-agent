package main

import (
	"context"
	"log"
	"time"

	"cmdb-agent/collector"
	pb "cmdb-agent/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("10.0.0.68:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
			panic(err)
		}
		log.Println("return code:", ack.Code)
		// 等待下次采集
		time.Sleep(nextExecTime.Sub(time.Now()))
	}
}

package main

import (
	"context"
	"fmt"
	"grpc2/message"
	"io"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8092", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)

	fmt.Println("客户端请求RPC调用: 双向流模式")
	orderIDs := []string{
		"201907300001",
		"201907310001",
		"201907310002",
	}

	orderInfoClient, err := orderServiceClient.GetOrderInfos(context.Background())
	for _, orderID := range orderIDs {
		orderRequest := message.OrderRequest{
			OrderId: orderID,
		}
		err := orderInfoClient.Send(&orderRequest)
		if err != nil {
			panic(err.Error())
		}
	}

	// 关闭
	orderInfoClient.CloseSend()

	for {
		orderInfo, err := orderInfoClient.Recv()
		if err == io.EOF {
			fmt.Println("读取结束")
			return
		}
		if err != nil {
			return
		}
		fmt.Println("读取到的信息: ", orderInfo)
	}
}

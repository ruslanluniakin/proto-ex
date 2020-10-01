package main

import (
	"context"
	"fmt"

	"github.com/ruslanluniakin/proto-ex/grpc-example/api/pingpongapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:55555", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	ppClient := pingpongapi.NewPingPongClient(conn)

	rsp, err := ppClient.PingPong(context.Background(), &pingpongapi.PingPongReq{
		Text: "ping",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("got response: %s\n", rsp.GetText())

	//for i := 0; i < 10; i++ {
	//	rsp, err := ppClient.PingPong(context.Background(), &pingpongapi.PingPongReq{
	//		Text: "ping",
	//	})
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("got response: %s\n", rsp.GetText())
	//}

}

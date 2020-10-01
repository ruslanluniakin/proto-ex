package main

import (
	"fmt"
	"net"

	"github.com/ruslanluniakin/proto-ex/api/pingpongapi"
	"github.com/ruslanluniakin/proto-ex/internal/pingpong"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const port = ":55555"

func main() {
	listener, err := net.Listen("tcp", port)

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	ppServer := pingpong.New()

	pingpongapi.RegisterPingPongServer(grpcServer, ppServer)

	fmt.Printf("server now serves on port: %s\n", port)
	fmt.Println("to stop the server press: CTRL+C")

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}

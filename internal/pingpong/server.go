package pingpong

import (
	"context"
	"fmt"

	"github.com/ruslanluniakin/proto-ex/api/pingpongapi"
)

// api is an implementation of pingpongapi.PingPongServer
type api struct {
}

func New() pingpongapi.PingPongServer {
	return &api{}
}

func (a api) PingPong(
	_ context.Context,
	req *pingpongapi.PingPongReq,
) (*pingpongapi.PingPongRsp, error) {
	fmt.Printf("got request: %s\n", req.GetText())
	return &pingpongapi.PingPongRsp{Text: "pong"}, nil
}

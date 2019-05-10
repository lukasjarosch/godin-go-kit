package grpc

import (
	"google.golang.org/grpc"
	"github.com/lukasjarosch/godin-go-kit/internal/endpoint"
	kitGrpc "github.com/go-kit/kit/transport/grpc"
	grpc2 "github.com/lukasjarosch/godin-go-kit/internal/grpc"
	"github.com/lukasjarosch/godin-go-kit/internal/example"
)

func NewClient(conn *grpc.ClientConn, opts ...kitGrpc.ClientOption) endpoint.Set {
	return endpoint.Set{
		HelloEndpoint:kitGrpc.NewClient(
			conn,
			"example",
			"Hello",
			grpc2.EncodeHelloRequest,
			grpc2.DecodeHelloResponse,
			example.HelloResponse{},
			opts...
		).Endpoint(),
	}
}
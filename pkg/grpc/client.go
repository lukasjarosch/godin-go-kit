package grpc

import (
	googleGrpc"google.golang.org/grpc"
	"github.com/lukasjarosch/godin-go-kit/internal/endpoint"
	kitGrpc "github.com/go-kit/kit/transport/grpc"
	grpc "github.com/lukasjarosch/godin-go-kit/internal/grpc"
	"github.com/lukasjarosch/godin-go-kit/internal/api"
)

func NewClient(conn *googleGrpc.ClientConn, opts ...kitGrpc.ClientOption) endpoint.Set {
	return endpoint.Set{
		HelloEndpoint:kitGrpc.NewClient(
			conn,
			"example.ExampleService",
			"Hello",
			grpc.EncodeHelloRequest,
			grpc.DecodeHelloResponse,
			api.HelloResponse{},
			opts...
		).Endpoint(),
	}
}
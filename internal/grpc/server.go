package grpc

import (

	kitGrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/lukasjarosch/godin-go-kit/internal/endpoint"
	"github.com/lukasjarosch/godin-go-kit/internal/example"
	"github.com/go-kit/kit/log"
	"context"
)

// TODO: Use godin to generate the grpc.Handlers and the whole server
// Godin certainly needs something like: godin add endpoint

type grpcServer struct {
	hello kitGrpc.Handler
}

func NewGrpcServer(endpoints endpoint.Set, logger log.Logger) example.ExampleServiceServer {
	options := []kitGrpc.ServerOption{
		kitGrpc.ServerErrorLogger(logger),
	}

	return &grpcServer{
		hello:kitGrpc.NewServer(
			endpoints.HelloEndpoint,
			DecodeHelloRequest,
			EncodeHelloResponse,
			append(options)...,
		),
	}
}

func (g *grpcServer) Hello(ctx context.Context, request *example.HelloRequest) (*example.HelloResponse, error) {
	_, response, err := g.hello.ServeGRPC(ctx, request)
	if err != nil {
	    return nil, EncodeError(err)
	}
	return response.(*example.HelloResponse), nil
}
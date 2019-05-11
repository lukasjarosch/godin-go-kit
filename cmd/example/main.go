package main

import (
	"github.com/go-kit/kit/log"
	"os"
	"github.com/lukasjarosch/godin-go-kit/internal/service"
	"github.com/lukasjarosch/godin-go-kit/internal/service/implementation"
	"github.com/lukasjarosch/godin-go-kit/internal/service/middleware"
	googleGrpc "google.golang.org/grpc"
	"net"
	"github.com/lukasjarosch/godin-go-kit/internal/api"
	"github.com/lukasjarosch/godin-go-kit/internal/service/endpoint"
	"github.com/lukasjarosch/godin-go-kit/internal/grpc"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewJSONLogger(os.Stderr)
		logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)
	}

	var svc service.Example
	svc = implementation.NewExampleService(logger)
	svc = middleware.NewLogMiddleware(logger)(svc)

	endpoints := endpoint.Endpoints(svc)
	grpcHandler := grpc.NewGrpcServer(endpoints, logger)

	grpcServer := googleGrpc.NewServer()
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}

	logger.Log("transport", "gRPC", "addr", ":50051")
	api.RegisterExampleServiceServer(grpcServer, grpcHandler)
	if err := grpcServer.Serve(grpcListener); err != nil {
		panic(err)
	}

}

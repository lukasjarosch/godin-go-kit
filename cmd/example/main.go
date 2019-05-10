package main

import (
	"github.com/go-kit/kit/log"
	"os"
	"github.com/lukasjarosch/godin-go-kit/internal/service"
	"github.com/lukasjarosch/godin-go-kit/internal/middleware"
	googleGrpc "google.golang.org/grpc"
	"net"
	"github.com/lukasjarosch/godin-go-kit/internal/example"
	"github.com/lukasjarosch/godin-go-kit/internal/endpoint"
	"github.com/lukasjarosch/godin-go-kit/internal/grpc"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewJSONLogger(os.Stderr)
		logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var svc service.ExampleService
	svc = service.NewExampleService()
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
	example.RegisterExampleServiceServer(grpcServer, grpcHandler)
	if err := grpcServer.Serve(grpcListener); err != nil {
		panic(err)
	}

}

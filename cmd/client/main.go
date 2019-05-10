package main

import (
	"context"
	"fmt"
	"os"

	example "github.com/lukasjarosch/godin-go-kit/pkg/grpc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error :%v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := example.NewClient(conn)

	resp, err := client.Hello(context.Background(), "Hans")
	if err != nil {
		fmt.Printf("error :%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("response: %v\n", resp)
}

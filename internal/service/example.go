package service

import (
	"context"
)

// This file provides the actual implementation of the endpoints which make up the service
// The interfaces are defined in 'service.go'
type exampleService struct {
}

func NewExampleService() *exampleService {
	return &exampleService{}
}

func (s *exampleService) Hello(ctx context.Context, name string) (greeting string, err error) {
	return "Foo", nil
}
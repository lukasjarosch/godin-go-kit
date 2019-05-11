package example

import (
	"context"
	"github.com/go-kit/kit/log"
)

// This file provides the actual implementation of the endpoints which make up the service
// The interfaces are defined in 'service.go'
type exampleService struct {
	logger log.Logger
}

func NewExampleService(logger log.Logger) *exampleService {
	return &exampleService{logger:logger}
}

func (s *exampleService) Hello(ctx context.Context, name string) (greeting string, err error) {
	logger := log.With(s.logger, "endpoint", "Hello")

	logger.Log("msg", "done greeting")
	//return "Foo", nil
	return "" , ErrNotImplemented
}

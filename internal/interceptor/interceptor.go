package interceptor

import (
	"github.com/lukasjarosch/godin-go-kit/internal/service"
)

type Interceptor func(service service.ExampleService) service.ExampleService
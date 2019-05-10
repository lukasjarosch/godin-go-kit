package middleware

import (
	"github.com/lukasjarosch/godin-go-kit/internal/service"
)

type Middleware func(service service.ExampleService) service.ExampleService
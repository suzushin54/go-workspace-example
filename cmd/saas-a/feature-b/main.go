package main

import (
	"context"

	"github.com/suzushin54/go-workspace-example/pkg/logger"
)

func main() {
	LogLevel := "info"
	l := logger.CreateLogger(&LogLevel)
	l.Infof(context.Background(), "Hello from Feature B!")
}

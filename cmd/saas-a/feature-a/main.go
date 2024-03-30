package main

import (
	"context"
	"fmt"

	"github.com/suzushin54/go-workspace-example/cmd/saas-a/common/api-client"
	"github.com/suzushin54/go-workspace-example/pkg/logger"
)

func main() {
	LogLevel := "info"
	l := logger.CreateLogger(&LogLevel)
	l.Infof(context.Background(), "Hello from Feature A!")

	client := api_client.NewInMemoryClient()
	data, err := client.FetchData()
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	fmt.Println("Fetched data:", data)
}

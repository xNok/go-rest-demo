package main

import (
	"context"
	"fmt"
	"github.com/xNok/go-rest-demo/pkg/gin"
	"github.com/xNok/go-rest-demo/pkg/rest"
	"os"
)

func main() {
	s := gin.NewServer()
	if err := run(context.Background(), s); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, s rest.Server) error {
	return s.Run(8080)
}

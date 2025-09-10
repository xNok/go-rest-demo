package main

import (
	"context"
	"fmt"
	"github.com/xNok/go-rest-demo/pkg/rest"
	"github.com/xNok/go-rest-demo/pkg/standardlib"
	"os"
)

func main() {
	s := standardlib.NewServer()
	if err := run(context.Background(), s); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, s rest.Server) error {
	return s.Run(8080)
}

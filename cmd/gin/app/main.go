package main

import (
	"fmt"
	"github.com/xNok/go-rest-demo/pkg/gin"
	"os"
)

func main() {
	s := gin.NewServer()
	port := 8080
	if err := s.Run(port); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

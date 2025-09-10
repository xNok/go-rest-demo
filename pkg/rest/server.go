package rest

import "context"

type Server interface {
	Run(port int) error
	Stop(ctx context.Context) error
}

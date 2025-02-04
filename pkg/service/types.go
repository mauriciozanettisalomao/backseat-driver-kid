package service

import "context"

// Handler defines the behavior of the main method
type Handler interface {
	Run(ctx context.Context) error
}

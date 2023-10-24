package server

import "context"

type HttpInstance interface {
	Hello(ctx context.Context) (int, string, error)
	Echo(ctx context.Context, in string) (int, string, error)
}

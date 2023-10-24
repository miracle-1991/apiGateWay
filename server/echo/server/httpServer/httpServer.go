package httpServer

import (
	"context"
	"github.com/miracle-1991/apiGateWay/server/echo/config"
)

type IMPL struct{}

func (i *IMPL) Hello(ctx context.Context) (int, string, error) {
	return config.OK, "hello world", nil
}

func (i *IMPL) Echo(ctx context.Context, in string) (int, string, error) {
	return config.OK, in, nil
}

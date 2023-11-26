package service

import (
	"context"
	"github.com/miracle-1991/apiGateWay/server/echo/config"
	"github.com/miracle-1991/apiGateWay/server/echo/observable/trace"
)

type IMPL struct{}

func (i *IMPL) Hello(ctx context.Context) (int, string, error) {
	ctx, span := trace.Tracer.Start(ctx, "hello")
	defer span.End()

	return config.OK, "hello world", nil
}

func (i *IMPL) Echo(ctx context.Context, in string) (int, string, error) {
	ctx, span := trace.Tracer.Start(ctx, "echo")
	defer span.End()

	return config.OK, in, nil
}

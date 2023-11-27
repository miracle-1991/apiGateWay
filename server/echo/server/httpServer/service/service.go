package service

import (
	"context"
	"github.com/miracle-1991/apiGateWay/server/echo/config"
	"github.com/miracle-1991/apiGateWay/server/echo/observable/trace"
	echo "github.com/miracle-1991/apiGateWay/server/echo/proto"
	"github.com/miracle-1991/apiGateWay/server/echo/server/utils"
)

type IMPL struct{}

func (i *IMPL) Hello(ctx context.Context) (int, string, error) {
	ctx, span := trace.Tracer.Start(ctx, "http-service-hello")
	defer span.End()

	return config.OK, "hello world", nil
}

func (i *IMPL) Echo(ctx context.Context, in string) (int, string, error) {
	ctx, span := trace.Tracer.Start(ctx, "http-service-echo")
	defer span.End()

	return config.OK, in, nil
}

func (i *IMPL) FillGeoHash(ctx context.Context, request *echo.FillGeoHashRequest) (*echo.FillGeoHashResponse, error) {
	ctx, span := trace.Tracer.Start(ctx, "http-service-fillgeohash")
	defer span.End()
	hashes, err := utils.FillGeoHash(ctx, request)
	if err != nil {
		return nil, err
	}
	return hashes, nil
}

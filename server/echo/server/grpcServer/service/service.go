package service

import (
	"context"
	echo "github.com/miracle-1991/apiGateWay/server/echo/proto"
	"github.com/miracle-1991/apiGateWay/server/echo/server/utils"
)

type Service interface {
	FillGeoHash(context.Context, *echo.FillGeoHashRequest) (*echo.FillGeoHashResponse, error)
}

type GeosService struct{}

func (g *GeosService) FillGeoHash(ctx context.Context, request *echo.FillGeoHashRequest) (*echo.FillGeoHashResponse, error) {
	return utils.FillGeoHash(ctx, request)
}

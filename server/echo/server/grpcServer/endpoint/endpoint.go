package endpoint

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	echo "github.com/miracle-1991/apiGateWay/server/echo/proto"
	"github.com/miracle-1991/apiGateWay/server/echo/server/grpcServer/service"
	"strings"
)

type EchoEndPoints struct {
	EchoEndPoint endpoint.Endpoint
}

func (e EchoEndPoints) FillGeoHash(ctx context.Context, r *echo.FillGeoHashRequest) (*echo.FillGeoHashResponse, error) {
	resp, err := e.EchoEndPoint(ctx, r)
	return resp.(*echo.FillGeoHashResponse), err
}

type EchoRequest struct {
	RequestType string                   `json:"request_type"`
	Request     *echo.FillGeoHashRequest `json:"request"`
}

type EchoResponse struct {
	Result *echo.FillGeoHashResponse `json:"result"`
	Error  error                     `json:"error"`
}

func MakeEchoEndPoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(EchoRequest)
		var (
			res *echo.FillGeoHashResponse
		)

		if strings.EqualFold(req.RequestType, "FillGeoHash") {
			res, err = svc.FillGeoHash(ctx, req.Request)
		} else {
			return nil, errors.New("NotSupport")
		}

		return EchoResponse{res, err}, nil
	}
}

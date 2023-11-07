package service

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	echo "github.com/miracle-1991/apiGateWay/server/echo/proto"
	"github.com/miracle-1991/apiGateWay/server/echo/server/grpcServer/endpoint"
)

type grpcServer struct {
	echo.UnimplementedGeoServiceServer
	FillGeoHashHandler grpc.Handler
}

func (g *grpcServer) FillGeoHash(ctx context.Context, r *echo.FillGeoHashRequest) (*echo.FillGeoHashResponse, error) {
	_, resp, err := g.FillGeoHashHandler.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*echo.FillGeoHashResponse), nil
}

func NewEchoServer(points endpoint.EchoEndPoints) echo.GeoServiceServer {
	return &grpcServer{
		FillGeoHashHandler: grpc.NewServer(
			points.EchoEndPoint,
			DecodeRequest,
			EncodeResponse,
		),
	}
}

func DecodeRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*echo.FillGeoHashRequest)
	return endpoint.EchoRequest{
		RequestType: "FillGeoHash",
		Request:     req,
	}, nil
}

func EncodeResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.EchoResponse)
	return resp.Result, resp.Error
}

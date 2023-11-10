package grpcServer

import (
	"fmt"
	"github.com/miracle-1991/apiGateWay/server/echo/config"
	echo "github.com/miracle-1991/apiGateWay/server/echo/proto"
	"github.com/miracle-1991/apiGateWay/server/echo/server/grpcServer/endpoint"
	"github.com/miracle-1991/apiGateWay/server/echo/server/grpcServer/service"
	myTransPort "github.com/miracle-1991/apiGateWay/server/echo/server/grpcServer/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func StartGrpcServer() {
	var svc service.Service
	svc = &service.GeosService{}

	endPoint := endpoint.MakeEchoEndPoint(svc)
	endPoints := endpoint.EchoEndPoints{
		EchoEndPoint: endPoint,
	}
	handler := myTransPort.NewEchoServer(endPoints)

	address := fmt.Sprintf(":%d", config.GRPC_PORT)
	ls, _ := net.Listen("tcp", address)
	gRPCServer := grpc.NewServer()

	healthServer := health.NewServer()
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(gRPCServer, healthServer)
	echo.RegisterGeoServiceServer(gRPCServer, handler)
	gRPCServer.Serve(ls)
}

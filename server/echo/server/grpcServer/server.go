package grpcServer

import (
	"fmt"
	"github.com/miracle-1991/apiGateWay/server/echo/config"
	echo "github.com/miracle-1991/apiGateWay/server/echo/proto"
	"github.com/miracle-1991/apiGateWay/server/echo/server/grpcServer/endpoint"
	"github.com/miracle-1991/apiGateWay/server/echo/server/grpcServer/service"
	myTransPort "github.com/miracle-1991/apiGateWay/server/echo/server/grpcServer/transport"
	"google.golang.org/grpc"
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
	echo.RegisterGeoServiceServer(gRPCServer, handler)

	gRPCServer.Serve(ls)
}

package main

import (
	"fmt"
	"github.com/miracle-1991/apiGateWay/server/echo/config"
	"github.com/miracle-1991/apiGateWay/server/echo/observable/trace"
	"github.com/miracle-1991/apiGateWay/server/echo/register"
	"github.com/miracle-1991/apiGateWay/server/echo/server/grpcServer"
	"github.com/miracle-1991/apiGateWay/server/echo/server/httpServer"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	//register consul
	c, err := register.NewConsul(config.CONSUL_ADDR)
	if err != nil {
		panic("failed to create consul client: " + err.Error())
	}
	ipOBJ, _ := register.GetOutboundIP()
	ip := ipOBJ.String()

	tags := []string{"version:" + strconv.Itoa(config.VER)}
	httpServiceID, grpcServiceID, err := c.RegisterService(config.SERVICENAME, ip, config.HTTP_PORT, config.GRPC_PORT, tags)
	if err != nil {
		panic("failed to register to consul: " + err.Error())
	}
	fmt.Printf("success register to consul, httpServiceID: %s, grpcServiceID: %s\n", httpServiceID, grpcServiceID)

	// register tracer
	err = trace.Register()
	if err != nil {
		fmt.Printf("failed to register trace, err:%v\n", err)
		panic(err)
	} else {
		fmt.Printf("register trace success\n")
	}

	// start http server
	go func() {
		httpServer.StartHttpServer()
	}()

	// start grpc server
	go func() {
		grpcServer.StartGrpcServer()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	_ = c.Deregister(httpServiceID)
	_ = c.Deregister(grpcServiceID)
	_ = trace.UnRegister()
}

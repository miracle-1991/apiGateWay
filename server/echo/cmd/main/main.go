package main

import (
	"fmt"
	"github.com/miracle-1991/apiGateWay/server/echo/config"
	"github.com/miracle-1991/apiGateWay/server/echo/register"
	"github.com/miracle-1991/apiGateWay/server/echo/router"
	"github.com/miracle-1991/apiGateWay/server/echo/server/grpcServer"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	route := router.Init()
	endPoint := fmt.Sprintf(":%d", config.HTTP_PORT)

	//register
	c, err := register.NewConsul(config.CONSUL_ADDR)
	if err != nil {
		panic("failed to create consul client: " + err.Error())
	}
	ipOBJ, _ := register.GetOutboundIP()
	ip := ipOBJ.String()
	serviceID := fmt.Sprintf("%s-%s-%d", config.SERVICENAME, ip, config.HTTP_PORT)
	tags := []string{"version:" + strconv.Itoa(config.VER)}
	err = c.RegisterService(config.SERVICENAME, ip, config.HTTP_PORT, tags)
	if err != nil {
		panic("failed to register to consul: " + err.Error())
	}
	fmt.Printf("success register to consul, serviceID: %s\n", serviceID)

	// start http server
	go func() {
		server := &http.Server{
			Addr:    endPoint,
			Handler: route,
		}
		err = server.ListenAndServe()
		if err != nil {
			panic("failed to serve http: " + err.Error())
		}
	}()

	// start grpc server
	go func() {
		grpcServer.StartGrpcServer()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	err = c.Deregister(serviceID)
	if err == nil {
		fmt.Printf("success deregister to consul, serviceID: %s\n", serviceID)
	}
}

package main

import (
	"fmt"
	"github.com/miracle-1991/apiGateWay/server/echo/config"
	"github.com/miracle-1991/apiGateWay/server/echo/router"
	"net/http"
)

func main() {
	route := router.Init()
	endPoint := fmt.Sprintf(":%d", config.PORT)
	server := &http.Server{
		Addr:    endPoint,
		Handler: route,
	}
	_ = server.ListenAndServe()
}

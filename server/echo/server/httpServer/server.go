package httpServer

import (
	"fmt"
	"github.com/miracle-1991/apiGateWay/server/echo/config"
	"github.com/miracle-1991/apiGateWay/server/echo/server/httpServer/transport"
	"net/http"
)

func StartHttpServer() {
	route := transport.Init()
	endPoint := fmt.Sprintf(":%d", config.HTTP_PORT)
	server := &http.Server{
		Addr:    endPoint,
		Handler: route,
	}
	_ = server.ListenAndServe()
}

package transport

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/miracle-1991/apiGateWay/server/echo/config"
	"github.com/miracle-1991/apiGateWay/server/echo/observable/trace"
	echo "github.com/miracle-1991/apiGateWay/server/echo/proto"
	"github.com/miracle-1991/apiGateWay/server/echo/server/httpServer/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"io"
)

func helloHandler(c *gin.Context) {
	ctx, span := trace.Tracer.Start(c.Request.Context(), "transport-hello")
	defer span.End()

	impl := &service.IMPL{}
	code, resp, err := impl.Hello(ctx)
	if err != nil {
		c.JSON(code, gin.H{
			"version": config.VER,
			"error":   err.Error(),
		})
	} else {
		c.JSON(code, gin.H{
			"version": config.VER,
			"message": resp,
		})
	}
}

func echoHandler(c *gin.Context) {
	ctx, span := trace.Tracer.Start(c.Request.Context(), "transport-echo")
	defer span.End()

	impl := &service.IMPL{}
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(config.InvalidInput, gin.H{
			"version": config.VER,
			"error":   err.Error(),
		})
		return
	}

	code, resp, err := impl.Echo(ctx, string(bodyBytes))
	if err != nil {
		c.JSON(code, gin.H{
			"version": config.VER,
			"error":   err.Error(),
		})
	} else {
		c.JSON(code, gin.H{
			"version": config.VER,
			"message": resp,
		})
	}
}

func healthHandler(c *gin.Context) {
	c.JSON(config.OK, gin.H{})
}

func hashFillHandler(c *gin.Context) {
	ctx, span := trace.Tracer.Start(c.Request.Context(), "transport-hashfill")
	defer span.End()

	request := &echo.FillGeoHashRequest{}
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Data format is not valid",
		})
		return
	}

	impl := &service.IMPL{}
	hashes, err := impl.FillGeoHash(ctx, request)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed, error: %v", err),
		})
		return
	}
	c.JSON(200, hashes)
}

func Init() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(otelgin.Middleware(config.SERVICENAME))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.GET("/health", healthHandler)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/hello", helloHandler)
	r.POST("/echo", echoHandler)
	r.POST("/hashfill", hashFillHandler)
	return r
}

package transport

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/miracle-1991/apiGateWay/server/echo/config"
	"github.com/miracle-1991/apiGateWay/server/echo/server/httpServer/service"
	"io"
)

func hello(c *gin.Context) {
	impl := &service.IMPL{}
	code, resp, err := impl.Hello(c.Request.Context())
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

func echo(c *gin.Context) {
	impl := &service.IMPL{}
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(config.InvalidInput, gin.H{
			"version": config.VER,
			"error":   err.Error(),
		})
		return
	}

	code, resp, err := impl.Echo(c.Request.Context(), string(bodyBytes))
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

func health(c *gin.Context) {
	c.JSON(config.OK, gin.H{})
}

func Init() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.GET("/health", health)
	r.GET("/hello", hello)
	r.POST("/echo", echo)
	return r
}

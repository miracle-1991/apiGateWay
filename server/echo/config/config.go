package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	SERVICENAME                           = "echo"
	ENV_PORT                              = "PORT"
	ENV_VERSION                           = "VERSION"
	ENV_CONSUL_ADDR                       = "CONSUL_HTTP_ADDR"
	ENV_HEALTH_CHECK_TIMEOUT              = "HEALTH_CHECK_TIMEOUT"
	ENV_HEALTH_CHECK_INTERVAL             = "HEALTH_CHECK_INTERVAL"
	ENV_DEREGISTER_CRITICAL_SERVICE_AFTER = "DEREGISTER_CRITICAL_SERVICE_AFTER"
)

var PORT, VER int
var CONSUL_ADDR, HEALTH_CHECK_TIMEOUT, HEALTH_CHECK_INTERVAL, DEREGISTER_CRITICAL_SERVICE_AFTER string

func init() {
	port := os.Getenv(ENV_PORT)
	version := os.Getenv(ENV_VERSION)
	if port == "" || version == "" {
		panic("invalid port and version")
	}
	var err error
	PORT, err = strconv.Atoi(port)
	if err != nil {
		panic(fmt.Sprintf("parse port err: %v", err))
	}
	fmt.Printf("port: %v\n", PORT)

	VER, err = strconv.Atoi(version)
	if err != nil {
		panic(fmt.Sprintf("parse version error: %v", err))
	}
	fmt.Printf("version: %v\n", VER)

	CONSUL_ADDR = os.Getenv(ENV_CONSUL_ADDR)
	if CONSUL_ADDR == "" {
		panic("invalid consul addr")
	}
	fmt.Printf("consul addr: %v\n", CONSUL_ADDR)

	// 超过10秒不返回认为不健康
	HEALTH_CHECK_TIMEOUT = os.Getenv(ENV_HEALTH_CHECK_TIMEOUT)
	if HEALTH_CHECK_TIMEOUT == "" {
		HEALTH_CHECK_TIMEOUT = "10s"
	}
	// 每10秒检查一次健康
	HEALTH_CHECK_INTERVAL = os.Getenv(ENV_HEALTH_CHECK_INTERVAL)
	if HEALTH_CHECK_INTERVAL == "" {
		HEALTH_CHECK_INTERVAL = "1s"
	}

	// 超过1分钟后注销不健康的节点
	DEREGISTER_CRITICAL_SERVICE_AFTER = os.Getenv(ENV_DEREGISTER_CRITICAL_SERVICE_AFTER)
	if DEREGISTER_CRITICAL_SERVICE_AFTER == "" {
		DEREGISTER_CRITICAL_SERVICE_AFTER = "1m"
	}
}

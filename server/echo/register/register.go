package register

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/miracle-1991/apiGateWay/server/echo/config"
	"net"
)

type consul struct {
	client *api.Client
}

func NewConsul(addr string) (*consul, error) {
	cfg := api.DefaultConfig()
	cfg.Address = addr
	c, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &consul{c}, nil
}

func (c *consul) RegisterService(serviceName string, ip string, httpPort, grpcPort int, tags []string) (string, string, error) {
	// register http
	httpServiceID := fmt.Sprintf("%s-%s-%d", serviceName, ip, httpPort)
	httpSrv := &api.AgentServiceRegistration{
		ID:      httpServiceID,
		Name:    serviceName + "_http",
		Tags:    tags,
		Address: ip,
		Port:    httpPort,
		Check: &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d/health", ip, httpPort),
			Timeout:                        config.HEALTH_CHECK_TIMEOUT,
			Interval:                       config.HEALTH_CHECK_INTERVAL,
			DeregisterCriticalServiceAfter: config.DEREGISTER_CRITICAL_SERVICE_AFTER,
		},
	}
	err := c.client.Agent().ServiceRegister(httpSrv)
	if err != nil {
		fmt.Printf("failed to register http server, error: %v", err)
		return "", "", err
	}

	// register grpc
	grpcServiceID := fmt.Sprintf("%s-%s-%d", serviceName, ip, grpcPort)
	grpcSrv := &api.AgentServiceRegistration{
		ID:      grpcServiceID,
		Name:    serviceName + "_grpc",
		Tags:    tags,
		Address: ip,
		Port:    grpcPort,
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d", ip, grpcPort),
			Timeout:                        config.HEALTH_CHECK_TIMEOUT,
			Interval:                       config.HEALTH_CHECK_INTERVAL,
			DeregisterCriticalServiceAfter: config.DEREGISTER_CRITICAL_SERVICE_AFTER,
		},
	}
	err = c.client.Agent().ServiceRegister(grpcSrv)
	if err != nil {
		fmt.Printf("failed to register grpc server, error: %v", err)
		return "", "", err
	}
	return httpServiceID, grpcServiceID, nil
}

func (c *consul) Deregister(serviceID string) error {
	return c.client.Agent().ServiceDeregister(serviceID)
}

func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

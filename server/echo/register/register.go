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

func (c *consul) RegisterService(serviceName string, ip string, port int, tags []string) error {
	//health check
	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port),
		Name:    serviceName,
		Tags:    tags,
		Address: ip,
		Port:    port,
		Check: &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d/health", ip, port),
			Timeout:                        config.HEALTH_CHECK_TIMEOUT,
			Interval:                       config.HEALTH_CHECK_INTERVAL,
			DeregisterCriticalServiceAfter: config.DEREGISTER_CRITICAL_SERVICE_AFTER,
		},
	}
	return c.client.Agent().ServiceRegister(srv)
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

package service

import (
	consul "github.com/hashicorp/consul/api"
	"time"
)

type Service struct {
	Name        string
	TTL         time.Duration
	ConsulAgent *consul.Agent
}

func New(ttl time.Duration) (*Service, error) {
	s := new(Service)
	s.Name = "my-service"
	s.TTL = ttl

	ok, err := s.Check()
	if !ok {
		return nil, err
	}

	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return nil, err
	}
	s.ConsulAgent = c.Agent()

	serviceDef := &consul.AgentServiceRegistration{
		Name: s.Name,
		Check: &consul.AgentServiceCheck{
			TTL: s.TTL.String(),
		},
	}

	if err := s.ConsulAgent.ServiceRegister(serviceDef); err != nil {
		return nil, err
	}
	go s.UpdateTTL(s.Check)

	return s, nil
}

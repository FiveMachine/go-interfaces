package common

import (
	"context"
)

type CommonConfig struct {
	Disabled bool
	Port     int
	Root     string
}

type ServiceConfig struct {
	CommonConfig
	Ctx        context.Context
	Others     map[string]int
	PublicKey  []byte
	PrivateKey []byte
	SwarmKey   []byte
}

type SimpleConfig struct {
	CommonConfig
	Clients map[string]ClientConfig
}

func (c *ServiceConfig) Clone() *ServiceConfig {
	clone := &ServiceConfig{
		CommonConfig: c.CommonConfig,
		Ctx:          c.Ctx,
		Others:       make(map[string]int, 0),
		PrivateKey:   c.PrivateKey,
		PublicKey:    c.PublicKey,
		SwarmKey:     c.SwarmKey,
	}

	for key, value := range c.Others {
		clone.Others[key] = value
	}

	return clone
}

type ClientConfig struct {
	CommonConfig
}

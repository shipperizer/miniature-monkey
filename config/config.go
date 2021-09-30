package config

import (
	"go.uber.org/zap"

	"github.com/shipperizer/miniature-monkey/monitoring"
	"github.com/shipperizer/miniature-monkey/utils"
)

type CORSConfig struct {
	origins []string
}

func (c *CORSConfig) GetOrigins() []string {
	return c.origins
}

func NewCORSConfig(origins ...string) CORSConfigInterface {
	c := new(CORSConfig)
	c.origins = origins

	return c
}

type APIConfig struct {
	name string

	cors CORSConfigInterface

	monitor monitoring.MonitorInterface
	logger  *zap.SugaredLogger
}

func (c *APIConfig) GetServiceName() string {
	return c.name
}

func (c *APIConfig) GetMonitor() monitoring.MonitorInterface {
	return c.monitor
}

func (c *APIConfig) GetCORSConfig() CORSConfigInterface {
	return c.cors
}

func (c *APIConfig) GetLogger() *zap.SugaredLogger {
	return c.logger
}

func NewAPIConfig(name string, cors CORSConfigInterface, monitor monitoring.MonitorInterface, logger *zap.SugaredLogger) APIConfigInterface {
	c := new(APIConfig)
	c.name = name
	c.cors = cors
	c.monitor = monitor
	c.logger = logger

	if c.logger == nil {
		c.logger = utils.NewLogger("error")
	}

	return c
}

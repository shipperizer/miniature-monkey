package config

import (
	"go.uber.org/zap"

	"github.com/shipperizer/miniature-monkey/monitoring"
	"github.com/shipperizer/miniature-monkey/utils"
)

// CORSConfig holds the origins to be CORS-allowed
type CORSConfig struct {
	origins []string
}

// GetOrigins returns a list of allowed origins
func (c *CORSConfig) GetOrigins() []string {
	return c.origins
}

// NewCORSConfig is the builder method to get a new CORSConfig
func NewCORSConfig(origins ...string) CORSConfigInterface {
	c := new(CORSConfig)
	c.origins = origins

	return c
}

// APIConfig is the main configurable object for the core.API object
// has attributes for CORS, the monitor object (see moniotring.MonitorInterface) adn the logger
type APIConfig struct {
	name string

	cors CORSConfigInterface

	monitor monitoring.MonitorInterface
	logger  *zap.SugaredLogger
}

// GetServiceName returns the a friendly name for the service
func (c *APIConfig) GetServiceName() string {
	return c.name
}

// GetMonitor returns the monitor object
func (c *APIConfig) GetMonitor() monitoring.MonitorInterface {
	return c.monitor
}

// GetCORSConfig returns the CORS config
func (c *APIConfig) GetCORSConfig() CORSConfigInterface {
	return c.cors
}

// GetLogger returns the logger
func (c *APIConfig) GetLogger() *zap.SugaredLogger {
	return c.logger
}

// NewAPIConfig returns a config object for the API object, if the logger arg is empty a new one with error level is created
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

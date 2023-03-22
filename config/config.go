package config

import (
	"github.com/shipperizer/miniature-monkey/v2/logging"
	core "github.com/shipperizer/miniature-monkey/v2/monitoring/core"
	trace "go.opentelemetry.io/otel/trace"
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
func NewCORSConfig(origins ...string) *CORSConfig {
	c := new(CORSConfig)
	c.origins = origins

	return c
}

// APIConfig is the main configurable object for the core.API object
// has attributes for CORS, the monitor object (see moniotring.MonitorInterface) adn the logger
type APIConfig struct {
	name string

	cors CORSConfigInterface

	tracer  trace.Tracer
	monitor core.MonitorInterface
	logger  logging.LoggerInterface
}

// GetServiceName returns the a friendly name for the service
func (c *APIConfig) GetServiceName() string {
	return c.name
}

// GetTracer returns the tracer object
func (c *APIConfig) GetTracer() trace.Tracer {
	return c.tracer
}

// GetMonitor returns the monitor object
func (c *APIConfig) GetMonitor() core.MonitorInterface {
	return c.monitor
}

// GetCORSConfig returns the CORS config
func (c *APIConfig) GetCORSConfig() CORSConfigInterface {
	return c.cors
}

// GetLogger returns the logger
func (c *APIConfig) GetLogger() logging.LoggerInterface {
	return c.logger
}

// NewAPIConfig returns a config object for the API object, if the logger arg is empty a new one with error level is created
func NewAPIConfig(name string, cors CORSConfigInterface, tracer trace.Tracer, monitor core.MonitorInterface, logger logging.LoggerInterface) *APIConfig {
	c := new(APIConfig)
	c.name = name
	c.cors = cors
	c.tracer = tracer
	c.monitor = monitor
	c.logger = logger

	if c.logger == nil {
		c.logger = logging.NewLogger("error")
	}

	return c
}

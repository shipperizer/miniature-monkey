package monitoring

import (
	"github.com/shipperizer/miniature-monkey/v2/logging"
	types "github.com/shipperizer/miniature-monkey/v2/monitoring/types"
	"go.uber.org/zap"
)

// MonitorConfig object
type MonitorConfig struct {
	metrics []types.MetricInterface
	service string
	logger  logging.LoggerInterface
}

func (c *MonitorConfig) GetMetrics() []types.MetricInterface {
	return c.metrics
}

func (c *MonitorConfig) GetService() string {
	return c.service
}

func (c *MonitorConfig) GetLogger() logging.LoggerInterface {
	return c.logger
}

func NewMonitorConfig(service string, metrics []types.MetricInterface, logger logging.LoggerInterface) *MonitorConfig {
	c := new(MonitorConfig)

	c.metrics = metrics
	c.service = service

	c.logger = logger

	if c.logger == nil {
		c.logger = zap.NewNop().Sugar()
	}

	return c
}

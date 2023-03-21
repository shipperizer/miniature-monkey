package monitoring

import (
	"github.com/shipperizer/miniature-monkey/v2/logging"
	types "github.com/shipperizer/miniature-monkey/v2/monitoring/types"
)

type MonitorConfigInterface interface {
	GetMetrics() []types.MetricInterface
	GetService() string
	GetLogger() logging.LoggerInterface
}

// MonitorInterface interface
// TODO @shipperizer bad pattern exposing internal types as interfaces, refactor
type MonitorInterface interface {
	GetService() string
	GetMetric(string) (types.MetricInterface, error)
	AddMetrics(...types.MetricInterface) error
	AddMetric(types.MetricInterface) error
}

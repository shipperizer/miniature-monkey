// Interface definitions
package monitoring

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

// MonitorInterface interface
type MonitorInterface interface {
	GetService() string
	GetMetric(metric string) *prometheus.GaugeVec
	Incr(metric string, opts map[string]string)
}

// MiddlewareInterface is the gorilla.mux Middleware interface
type MiddlewareInterface interface {
	APICount() mux.MiddlewareFunc
	APITime() mux.MiddlewareFunc
}

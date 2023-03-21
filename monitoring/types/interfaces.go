package monitoring

import "github.com/prometheus/client_golang/prometheus"

type MetricInterface interface {
	Inc(tags map[string]string) error
	Dec(tags map[string]string) error
	Add(value float64, tags map[string]string) error
	Sub(value float64, tags map[string]string) error
	Set(value float64, tags map[string]string) error
	Observe(value float64, tags map[string]string) error
	Collector() prometheus.Collector
	Name() string
}

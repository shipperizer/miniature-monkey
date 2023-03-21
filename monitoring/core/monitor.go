package monitoring

import (
	"fmt"

	"github.com/shipperizer/miniature-monkey/v2/logging"
	types "github.com/shipperizer/miniature-monkey/v2/monitoring/types"

	"github.com/prometheus/client_golang/prometheus"
)

// Monitor object
type Monitor struct {
	metrics map[string]types.MetricInterface

	service string

	logger logging.LoggerInterface
}

func (m *Monitor) GetService() string {
	return m.service
}

func (m *Monitor) GetLogger() logging.LoggerInterface {
	return m.logger
}

func (m *Monitor) GetMetric(metric string) (types.MetricInterface, error) {
	if m, ok := m.metrics[metric]; ok {
		return m, nil
	}

	m.logger.Debugf("No metric found with the name %s", metric)

	return nil, fmt.Errorf("no metric found with the name %s", metric)
}

// AddMetrics will try to register the metrics passed in, if already present it will Unregister them so that the newest added is the one used
// by definition from https://github.com/prometheus/client_golang/blob/main/prometheus/registry.go#L121
// Two Collectors are considered equal if their Describe method yields the same set of descriptors.
func (m *Monitor) AddMetrics(metrics ...types.MetricInterface) error {
	for _, metric := range metrics {
		if err := m.AddMetric(metric); err != nil {
			return err
		}
	}

	return nil
}

// AddMetric will try to register the metric passed in, if already present it will Unregister them so that the newest added is the one used
// by definition from https://github.com/prometheus/client_golang/blob/main/prometheus/registry.go#L121
// Two Collectors are considered equal if their Describe method yields the same set of descriptors.
func (m *Monitor) AddMetric(metric types.MetricInterface) error {
	if metric == nil {
		m.logger.Debug("metric is nil")
		return fmt.Errorf("metric is nil")
	}

	c := metric.Collector()

	if c == nil {
		m.logger.Debug("metric is not initialized")
		return fmt.Errorf("metric is not initialized")
	}

	if err := prometheus.Register(c); err != nil {
		m.logger.Debugf("%s: metric already registered with the name %v...unregistering", err, metric.Name())
		prometheus.Unregister(c)

		if err := prometheus.Register(c); err != nil {
			// not returning an error due to the fact we are using the shared prometheus register
			m.logger.Errorf("%s: second attempt went wrong for %v...metric is not registered, continue anyway", err, metric.Name())
		}
	}

	// if above is successful, add it to the metrics attribute
	m.metrics[metric.Name()] = metric

	return nil
}

// NewMonitor returns a MonitorInterface, an enhancement (transitioning) on MonitorService
func NewMonitor(cfg MonitorConfigInterface) *Monitor {
	m := new(Monitor)

	m.service = cfg.GetService()
	m.logger = cfg.GetLogger()
	m.metrics = make(map[string]types.MetricInterface)

	iMetrics := make([]types.MetricInterface, 0)

	for _, metric := range cfg.GetMetrics() {
		iMetrics = append(iMetrics, metric)
	}

	m.AddMetrics(iMetrics...)

	return m
}

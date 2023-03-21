package monitoring

import (
	"testing"

	"github.com/golang/mock/gomock"
	types "github.com/shipperizer/miniature-monkey/v2/monitoring/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

//go:generate mockgen -build_flags=--mod=mod -package monitoring -destination ./mock_types.go -source=../types/interfaces.go
//go:generate mockgen -build_flags=--mod=mod -package monitoring -destination ./mock_core.go -source=interfaces.go

func TestNewMonitorImplementsMonitorInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfig := NewMockMonitorConfigInterface(ctrl)
	mockMetric := NewMockMetricInterface(ctrl)

	mockConfig.EXPECT().GetService().Times(1).Return("test")
	mockConfig.EXPECT().GetLogger().Times(1).Return(zap.NewNop().Sugar())
	mockConfig.EXPECT().GetMetrics().Times(1).Return([]types.MetricInterface{mockMetric})

	mockMetric.EXPECT().Collector().AnyTimes().Return(nil)

	monitor := NewMonitor(mockConfig)

	assert := assert.New(t)
	assert.Implements((*MonitorInterface)(nil), monitor)
}

func TestMonitorAddMetrics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfig := NewMockMonitorConfigInterface(ctrl)
	mockMetric := NewMockMetricInterface(ctrl)

	mockConfig.EXPECT().GetService().Times(1).Return("test")
	mockConfig.EXPECT().GetLogger().Times(1).Return(zap.NewNop().Sugar())
	mockConfig.EXPECT().GetMetrics().Times(1).Return([]types.MetricInterface{mockMetric})

	mockMetric.EXPECT().Collector().AnyTimes().Return(nil)

	monitor := NewMonitor(mockConfig)

	metrics := []types.MetricInterface{
		types.NewMetric(types.HISTOGRAM, "h_test"),
		types.NewMetric(types.GAUGE, "g_test"),
		types.NewMetric(types.COUNTER, "c_test"),
	}

	err := monitor.AddMetrics(metrics...)

	assert := assert.New(t)

	assert.Nil(err)

	for _, metric := range metrics {
		m, err := monitor.GetMetric(metric.Name())

		assert.Nil(err)
		assert.NotNil(m)
	}

}

func TestMonitorAddMetricsDoNotErrorWhenMetricExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger, _ := zap.NewDevelopment()
	mockConfig := NewMockMonitorConfigInterface(ctrl)
	mockMetric := NewMockMetricInterface(ctrl)

	metrics := []types.MetricInterface{
		types.NewMetric(types.HISTOGRAM, "h_test"),
		types.NewMetric(types.GAUGE, "g_test"),
		types.NewMetric(types.COUNTER, "c_test"),
	}

	mockConfig.EXPECT().GetService().Times(1).Return("test")
	mockConfig.EXPECT().GetLogger().Times(1).Return(logger.Sugar())
	mockConfig.EXPECT().GetMetrics().Times(1).Return(metrics)
	mockMetric.EXPECT().Collector().AnyTimes().Return(nil)

	monitor := NewMonitor(mockConfig)

	assert := assert.New(t)
	assert.Nil(monitor.AddMetrics(metrics[0]))

	m, err := monitor.GetMetric(metrics[0].Name())
	assert.Nil(err)
	assert.NotNil(m)
}

func TestMonitorGetMetric(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger, _ := zap.NewDevelopment()
	mockConfig := NewMockMonitorConfigInterface(ctrl)

	metrics := []types.MetricInterface{
		types.NewMetric(types.HISTOGRAM, "h_test"),
		types.NewMetric(types.GAUGE, "g_test"),
		types.NewMetric(types.COUNTER, "c_test"),
	}

	mockConfig.EXPECT().GetService().Times(1).Return("test")
	mockConfig.EXPECT().GetLogger().Times(1).Return(logger.Sugar())
	mockConfig.EXPECT().GetMetrics().Times(1).Return(metrics)

	monitor := NewMonitor(mockConfig)

	assert := assert.New(t)

	m, err := monitor.GetMetric("h_test")
	assert.Nil(err)
	assert.NotNil(m)
}

func TestMonitorGetMetricErrorIFNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfig := NewMockMonitorConfigInterface(ctrl)

	metrics := []types.MetricInterface{
		types.NewMetric(types.HISTOGRAM, "h_test"),
		types.NewMetric(types.GAUGE, "g_test"),
		types.NewMetric(types.COUNTER, "c_test"),
	}

	mockConfig.EXPECT().GetService().Times(1).Return("test")
	mockConfig.EXPECT().GetLogger().Times(1).Return(zap.NewNop().Sugar())
	mockConfig.EXPECT().GetMetrics().Times(1).Return(metrics)

	monitor := NewMonitor(mockConfig)

	assert := assert.New(t)

	m, err := monitor.GetMetric("test")
	assert.NotNil(err)
	assert.Nil(m)
}

func TestMonitorGetLogger(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfig := NewMockMonitorConfigInterface(ctrl)
	mockMetric := NewMockMetricInterface(ctrl)
	logger := zap.NewNop().Sugar()
	metrics := []types.MetricInterface{mockMetric}

	mockConfig.EXPECT().GetService().Times(1).Return("test")
	mockConfig.EXPECT().GetLogger().Times(1).Return(logger)
	mockConfig.EXPECT().GetMetrics().Times(1).Return(metrics)
	mockMetric.EXPECT().Collector().AnyTimes().Return(nil)

	monitor := NewMonitor(mockConfig)

	assert := assert.New(t)
	assert.Equal(logger, monitor.GetLogger())

}

func TestMonitorGetService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfig := NewMockMonitorConfigInterface(ctrl)

	mockConfig.EXPECT().GetService().Times(1).Return("test")
	mockConfig.EXPECT().GetLogger().Times(1).Return(zap.NewNop().Sugar())
	mockConfig.EXPECT().GetMetrics().Times(1).Return(nil)

	monitor := NewMonitor(mockConfig)

	assert := assert.New(t)
	assert.Equal("test", monitor.GetService())

}

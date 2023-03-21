package monitoring

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	types "github.com/shipperizer/miniature-monkey/v2/monitoring/types"
)

//go:generate mockgen -build_flags=--mod=mod -package monitoring -destination ./mock_types.go -source=../types/interfaces.go

func TestMonitorConfigGetService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := NewMonitorConfig(
		"test",
		nil,
		nil,
	)

	assert := assert.New(t)

	assert.Equal("test", cfg.GetService(), "values should have matched")
}

func TestMonitorConfigGetMetrics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMetric := NewMockMetricInterface(ctrl)

	cfg := NewMonitorConfig(
		"test",
		[]types.MetricInterface{mockMetric},
		nil,
	)

	assert := assert.New(t)

	assert.Equal([]types.MetricInterface{mockMetric}, cfg.GetMetrics())
}

func TestMonitorConfigGetLogger(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := zap.NewNop().Sugar()

	cfg := NewMonitorConfig(
		"test",
		nil,
		logger,
	)

	assert := assert.New(t)

	assert.Equal(logger, cfg.GetLogger(), "values should have matched")
}

func TestMonitorConfigGetLoggerDefault(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := NewMonitorConfig(
		"test",
		nil,
		nil,
	)

	assert := assert.New(t)

	assert.NotNil(cfg.GetLogger())
}

package config

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shipperizer/miniature-monkey/v2/logging"
)

//go:generate mockgen -package config -destination ./mock_config.go -source=./interfaces.go CORSConfigInterface
//go:generate mockgen -package config -destination ./mock_monitor.go -source=../monitoring/core/interfaces.go MonitorInterface

func TestAPIConfigGetCORSConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCORS := NewMockCORSConfigInterface(ctrl)
	mockMonitor := NewMockMonitorInterface(ctrl)

	cfg := NewAPIConfig(
		"shipperizer",
		mockCORS,
		mockMonitor,
		nil,
	)

	assert := assert.New(t)

	assert.Equal(mockCORS, cfg.GetCORSConfig())
}

func TestAPIConfigGetService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCORS := NewMockCORSConfigInterface(ctrl)
	mockMonitor := NewMockMonitorInterface(ctrl)

	svc := "shipperizer"

	cfg := NewAPIConfig(
		svc,
		mockCORS,
		mockMonitor,
		nil,
	)

	assert := assert.New(t)

	assert.Equal(svc, cfg.GetServiceName())
}

func TestAPIConfigGetLogger(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCORS := NewMockCORSConfigInterface(ctrl)
	mockMonitor := NewMockMonitorInterface(ctrl)
	logger := logging.NewLogger("debug")

	cfg := NewAPIConfig(
		"shipperizer",
		mockCORS,
		mockMonitor,
		logger,
	)

	assert := assert.New(t)

	assert.Equal(logger, cfg.GetLogger())
}

func TestAPIConfigGetLoggerIfNilPassed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCORS := NewMockCORSConfigInterface(ctrl)
	mockMonitor := NewMockMonitorInterface(ctrl)

	cfg := NewAPIConfig(
		"shipperizer",
		mockCORS,
		mockMonitor,
		nil,
	)

	assert := assert.New(t)

	assert.Implements((*logging.LoggerInterface)(nil), cfg.GetLogger())
}

func TestAPIConfigGetMonitor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCORS := NewMockCORSConfigInterface(ctrl)
	mockMonitor := NewMockMonitorInterface(ctrl)

	cfg := NewAPIConfig(
		"shipperizer",
		mockCORS,
		mockMonitor,
		nil,
	)

	assert := assert.New(t)

	assert.Equal(mockMonitor, cfg.GetMonitor())
}

func TestCORSConfigGetOrigins(t *testing.T) {
	cfg := NewCORSConfig("shipperizer.com", "test.com")

	assert := assert.New(t)

	assert.Equal([]string{"shipperizer.com", "test.com"}, cfg.GetOrigins())
}

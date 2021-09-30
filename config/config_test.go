package config

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/shipperizer/miniature-monkey/utils"
)

//go:generate mockgen -package config -destination ./mock_config.go -source=./interfaces.go CORSConfigInterface
//go:generate mockgen -package config -destination ./mock_monitor.go -source=../monitoring/interfaces.go MonitorInterface

func TestAPIConfigImplementsAPIConfigInterface(t *testing.T) {
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

	assert.Implements((*APIConfigInterface)(nil), cfg, "object should have implemented APIConfigInterface")
}

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
	logger := utils.NewLogger("debug")

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

	assert.IsType(&zap.SugaredLogger{}, cfg.GetLogger())
	assert.True(cfg.GetLogger().Desugar().Core().Enabled(zapcore.ErrorLevel), "error level should be enabled")
	assert.True(cfg.GetLogger().Desugar().Core().Enabled(zapcore.WarnLevel), "warning level should not be enabled")
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

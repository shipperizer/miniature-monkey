package core

import (
	"net/http"

	"testing"

	chi "github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//go:generate mockgen -package core -destination ./mock_webiface.go -source=../webiface/interfaces.go APIInterface
//go:generate mockgen -package core -destination ./mock_config.go -source=../config/interfaces.go CORSConfigInterface
//go:generate mockgen -package core -destination ./mock_monitor.go -source=../monitoring/core/interfaces.go MonitorInterface
//go:generate mockgen -package core -destination ./mock_core.go -source=./interfaces.go APIConfigInterface

func TestAPIRouterReturnChiMux(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockConfig := NewMockAPIConfigInterface(ctrl)

	mockConfig.EXPECT().GetServiceName().Times(1).Return("test")
	mockConfig.EXPECT().GetMonitor().Times(1).Return(nil)
	mockConfig.EXPECT().GetLogger().Times(1).Return(nil)
	mockConfig.EXPECT().GetCORSConfig().Times(1).Return(nil)

	newApi := NewAPI(mockConfig)

	assert := assert.New(t)

	assert.IsType(new(chi.Mux), newApi.Router(), "object should have been of type *chi.Mux")
}

func TestAPIHandlerImplementHttpHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockConfig := NewMockAPIConfigInterface(ctrl)
	mockCORS := NewMockCORSConfigInterface(ctrl)

	mockCORS.EXPECT().GetOrigins().Times(1).Return([]string{"test.com"})
	mockConfig.EXPECT().GetServiceName().Times(1).Return("test")
	mockConfig.EXPECT().GetMonitor().Times(1).Return(nil)
	mockConfig.EXPECT().GetLogger().Times(1).Return(nil)
	mockConfig.EXPECT().GetCORSConfig().Times(1).Return(mockCORS)

	newApi := NewAPI(mockConfig)

	assert := assert.New(t)

	assert.Implements((*http.Handler)(nil), newApi.Handler(), "object should have implemented http.Handler")
}

func TestAPIAddStatusAndMonitoring(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockConfig := NewMockAPIConfigInterface(ctrl)
	mockMonitor := NewMockMonitorInterface(ctrl)

	mockConfig.EXPECT().GetCORSConfig().Times(1).Return(nil)
	mockConfig.EXPECT().GetServiceName().Times(1).Return("test")
	mockConfig.EXPECT().GetMonitor().Times(1).Return(mockMonitor)
	mockConfig.EXPECT().GetLogger().Times(1).Return(nil)
	mockMonitor.EXPECT().GetService().Times(1).Return("test")
	mockMonitor.EXPECT().AddMetrics(gomock.Any()).Times(1).Return(nil)

	newAPI := NewAPI(mockConfig)

	assert := assert.New(t)

	statusRoute := false
	metricsRoute := false

	for _, route := range newAPI.Router().Routes() {
		if route.Pattern == "/api/v1/status" {
			statusRoute = true
		} else if route.Pattern == "/api/v1/metrics" {
			metricsRoute = true
		}
	}

	assert.True(statusRoute, "object should have a route named v1.status")
	assert.True(metricsRoute, "object should have a route named v1.metrics")
}

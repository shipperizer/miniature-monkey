package core

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/shipperizer/miniature-monkey/webiface"

	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//go:generate mockgen -package core -destination ./mock_webiface.go -source=../webiface/interfaces.go APIInterface
//go:generate mockgen -package core -destination ./mock_config.go -source=../config/interfaces.go APIConfigInterface,CORSConfigInterface
//go:generate mockgen -package core -destination ./mock_monitor.go -source=../monitoring/interfaces.go MonitorInterface

func TestAPIImplementsAPIInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockConfig := NewMockAPIConfigInterface(ctrl)

	mockConfig.EXPECT().GetServiceName().Times(1).Return("test")
	mockConfig.EXPECT().GetMonitor().Times(1).Return(nil)
	mockConfig.EXPECT().GetLogger().Times(1).Return(nil)

	newApi := NewAPI(mockConfig)

	assertion := assert.New(t)

	assertion.Implements((*webiface.APIInterface)(nil), newApi, "object should have implemented APIInterface")
}

func TestAPIRouterReturnMuxRouter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockConfig := NewMockAPIConfigInterface(ctrl)

	mockConfig.EXPECT().GetServiceName().Times(1).Return("test")
	mockConfig.EXPECT().GetMonitor().Times(1).Return(nil)
	mockConfig.EXPECT().GetLogger().Times(1).Return(nil)

	newApi := NewAPI(mockConfig)

	assertion := assert.New(t)

	assertion.IsType(&mux.Router{}, newApi.Router(), "object should have been of type *mux.Router")
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

	assertion := assert.New(t)

	assertion.Implements((*http.Handler)(nil), newApi.Handler(), "object should have implemented http.Handler")
}

func TestAPIAddStatusAndMonitoring(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockConfig := NewMockAPIConfigInterface(ctrl)
	mockMonitor := NewMockMonitorInterface(ctrl)

	mockConfig.EXPECT().GetServiceName().Times(1).Return("test")
	mockConfig.EXPECT().GetMonitor().Times(1).Return(mockMonitor)
	mockConfig.EXPECT().GetLogger().Times(1).Return(nil)
	mockMonitor.EXPECT().GetService().Times(1).Return("test")

	newAPI := NewAPI(mockConfig)

	assertion := assert.New(t)

	assertion.NotNil(newAPI.Router().Get("v1.status"), "object should have a route named v1.status")
	assertion.NotNil(newAPI.Router().Get("v1.metrics"), "object should have a route named v1.metrics")
}

package monitoring

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

//go:generate mockgen -package monitoring -destination ./mock_monitor.go -source=./interfaces.go MonitorInterface

func TestMiddlewareAPICount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMonitor := NewMockMonitorInterface(ctrl)

	mockMonitor.EXPECT().Incr("http_api_count", gomock.Any()).Times(1)

	assert := assert.New(t)
	router := mux.NewRouter()
	NewBlueprint().Routes(router)
	router.Use(NewMiddleware(mockMonitor, "test").APICount())

	// setup metrics endpoint
	req, err := http.NewRequest(http.MethodGet, "/api/v1/metrics", nil)
	req.Header.Set("Content-Type", "application/json")
	assert.Equal(nil, err, "error should be nil")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
}

func TestMiddlewareAPITime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMonitor := NewMockMonitorInterface(ctrl)

	opts := prometheus.GaugeOpts{
		Name: "http_api_time",
		Help: "test http_api_time",
	}

	mockMonitor.EXPECT().GetMetric("http_api_time").Times(1).Return(prometheus.NewGaugeVec(opts, []string{}))

	assert := assert.New(t)
	router := mux.NewRouter()
	NewBlueprint().Routes(router)
	router.Use(NewMiddleware(mockMonitor, "test").APITime())

	// setup metrics endpoint
	req, err := http.NewRequest(http.MethodGet, "/api/v1/metrics", nil)
	req.Header.Set("Content-Type", "application/json")
	assert.Equal(nil, err, "error should be nil")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
}

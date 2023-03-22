package generics

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
)

//go:generate mockgen -build_flags=--mod=mod -package generics -destination ./mock_monitor.go -source=../../monitoring/core/interfaces.go MonitorInterface
//go:generate mockgen -build_flags=--mod=mod -package generics -destination ./mock_metrics.go -source=../../monitoring/types/interfaces.go MetricInterface
//go:generate mockgen -build_flags=--mod=mod -package generics -destination ./mock_logger.go -source=../../logging/interfaces.go LoggerInterface

type Blueprint struct{}

// Routes exposes the /api/v1/metrics endpoint for prometheus to scrape
func (b *Blueprint) Routes(router *chi.Mux) {
	router.Get("/api/v1/metrics", b.prometheusHTTP)
	router.Get("/api/test", b.test)
}

func (b *Blueprint) prometheusHTTP(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}

func (b *Blueprint) test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func TestMiddlewareAPITime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMonitor := NewMockMonitorInterface(ctrl)
	mockMetric := NewMockMetricInterface(ctrl)
	mockLogger := NewMockLoggerInterface(ctrl)
	mockMonitor.EXPECT().GetService().AnyTimes()
	mockMonitor.EXPECT().AddMetrics(gomock.Any()).AnyTimes()
	mockMonitor.EXPECT().GetMetric("http_server_handling_seconds_v1").Times(1).Return(mockMetric, nil)
	mockMetric.EXPECT().Observe(gomock.Any(), gomock.Any()).Times(1)

	assert := assert.New(t)

	router := chi.NewMux()

	router.Use(NewMiddleware(mockMonitor, mockLogger).APITime())

	new(Blueprint).Routes(router)

	// setup metrics endpoint
	req, err := http.NewRequest(http.MethodGet, "/api/test", nil)
	req.Header.Set("Content-Type", "application/json")
	assert.Nil(err, "error should be nil")

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
}

package generics

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	chi "github.com/go-chi/chi/v5"

	"github.com/shipperizer/miniature-monkey/v2/logging"
	core "github.com/shipperizer/miniature-monkey/v2/monitoring/core"
	types "github.com/shipperizer/miniature-monkey/v2/monitoring/types"
)

const (
	// IDPathRegex regexp used to swap the {id*} parameters in the path with simply id
	// supports alphabetic characters and underscores, no dashes
	IDPathRegex string = "{[a-zA-Z_]*}"
)

// Middleware is the monitoring middleware object implementing Prometheus monitoring
type Middleware struct {
	regex   *regexp.Regexp
	service string

	monitor core.MonitorInterface
	logger  logging.LoggerInterface
}

// APITime sets a labs_http_server_handling_seconds_v1 metric
func (mdw Middleware) APITime() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				path := chi.RouteContext(r.Context()).RoutePath
				tags := map[string]string{"service": mdw.service, "route": fmt.Sprintf("%s%s", r.Method, mdw.regex.ReplaceAll([]byte(path), []byte("id")))}

				m, err := mdw.monitor.GetMetric("http_server_handling_seconds_v1")

				startTime := time.Now()

				rec := &statusRecorder{w, http.StatusAccepted}

				next.ServeHTTP(rec, r)

				if err != nil {
					mdw.logger.Debugf("Error fetching metric: %s; keep going....", err)
				} else {
					// collect the status from the wrapper
					tags["status"] = fmt.Sprint(rec.status)

					m.Observe(time.Since(startTime).Seconds(), tags)
				}
			},
		)
	}
}

// custom prometheus metrics setup
// ###################################################################################
func (mdw *Middleware) registerMetrics() error {
	m := []types.MetricInterface{
		types.NewMetric(types.HISTOGRAM, "http_server_handling_seconds_v1", "service", "route", "status"),
	}

	return mdw.monitor.AddMetrics(m...)
}

// NewMiddleware returns a Middleware based on the type of monitor
func NewMiddleware(monitor core.MonitorInterface, logger logging.LoggerInterface) *Middleware {
	mdw := new(Middleware)

	mdw.monitor = monitor

	mdw.service = monitor.GetService()
	mdw.logger = logger

	mdw.regex = regexp.MustCompile(IDPathRegex)

	_ = mdw.registerMetrics()

	return mdw
}

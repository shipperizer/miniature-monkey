package monitoring

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// IDPathRegex regexp used to swap the {id*} parameters in the path with simply id
	// supports alphabetic characters and underscores, no dashes
	IDPathRegex string = "{[a-zA-Z_]*}"
)

// Middleware is the monitoring middleware object implementing Prometheus gateway
// monitoring
type Middleware struct {
	monitor MonitorInterface
	regex   *regexp.Regexp
	service string
}

// APICount increments a http_api_count metric on Prometheus
func (mdw Middleware) APICount() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path, _ := mux.CurrentRoute(r).GetPathTemplate()
			labels := map[string]string{"service": mdw.service, "route": fmt.Sprintf("%s%s", r.Method, mdw.regex.ReplaceAll([]byte(path), []byte("id")))}

			mdw.monitor.Incr("http_api_count", labels)

			next.ServeHTTP(w, r)
		})
	}
}

// APITime sets a http_api_time metric on Prometheus
func (mdw Middleware) APITime() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path, _ := mux.CurrentRoute(r).GetPathTemplate()
			labels := map[string]string{"service": mdw.service, "route": fmt.Sprintf("%s%s", r.Method, mdw.regex.ReplaceAll([]byte(path), []byte("id")))}

			gauge, err := mdw.monitor.GetMetric("http_api_time").GetMetricWith(labels)

			if err != nil {
				next.ServeHTTP(w, r)
			} else {
				timer := prometheus.NewTimer(prometheus.ObserverFunc(gauge.Set))
				next.ServeHTTP(w, r)
				timer.ObserveDuration()
			}
		})
	}
}

// NewMiddleware returns a Middleware based on the type of monitor
func NewMiddleware(monitor MonitorInterface, service string) MiddlewareInterface {
	return &Middleware{
		monitor: monitor,
		regex:   regexp.MustCompile(IDPathRegex),
		service: service,
	}

}

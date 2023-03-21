package monitoring

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Blueprint struct{}

// Routes exposes the /api/v1/metrics endpoint for prometheus to scrape
func (bp *Blueprint) Routes(router *chi.Mux) {
	router.Get("/api/v1/metrics", bp.prometheusHTTP)
}

func (b *Blueprint) prometheusHTTP(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}

// NewBlueprint returns a new initialized Blueprint object.
func NewBlueprint() *Blueprint {
	return new(Blueprint)
}

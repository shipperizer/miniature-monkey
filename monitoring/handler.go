package monitoring

import (
	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/shipperizer/miniature-monkey/webiface"
)

type Blueprint struct{}

// Routes exposes the /api/v1/metrics endpoint for prometheus to scrape
func (bp *Blueprint) Routes(router *mux.Router) {
	router.PathPrefix("/api/v1/metrics").Handler(promhttp.Handler()).Name("v1.metrics")
}

// NewBlueprint returns a new initialized Blueprint object.
func NewBlueprint() webiface.BlueprintInterface {
	return new(Blueprint)
}

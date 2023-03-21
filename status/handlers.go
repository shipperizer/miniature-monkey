// API handlers for the publish blueprint endpoints

package status

import (
	"encoding/json"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/shipperizer/miniature-monkey/v2/types"
)

type Blueprint struct{}

// Routes exposes the handler on a route attached to the router
func (bp *Blueprint) Routes(router *chi.Mux) {
	router.Get("/api/v1/status", bp.status)
}

// Status is a basic status endpoint returning ok
func (bp *Blueprint) status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := new(types.DataResponse)
	resp.Message = "ok"

	json.NewEncoder(w).Encode(resp)
}

// NewBlueprint returns a new initialized Blueprint object.
func NewBlueprint() *Blueprint {
	return new(Blueprint)
}

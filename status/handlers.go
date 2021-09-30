// API handlers for the publish blueprint endpoints

package status

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/shipperizer/miniature-monkey/types"
	"github.com/shipperizer/miniature-monkey/webiface"
)

type Blueprint struct{}

// Routes exposes the handler on a route attached to the router
func (bp *Blueprint) Routes(router *mux.Router) {
	router.HandleFunc("/api/v1/status", bp.Status).Methods(http.MethodGet).Name("v1.status")
}

// Status is a basic status endpoint returning ok
func (bp Blueprint) Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := new(types.DataResponse)
	resp.Message = "ok"

	json.NewEncoder(w).Encode(resp)
}

// NewBlueprint returns a new initialized Blueprint object.
func NewBlueprint() webiface.BlueprintInterface {
	return new(Blueprint)
}

// API handlers for the publish blueprint endpoints

package status

import (
	"encoding/json"
	"net/http"
	"runtime/debug"

	chi "github.com/go-chi/chi/v5"
	"github.com/shipperizer/miniature-monkey/v2/types"
)

type BuildInfo struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}

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

	if buildInfo := bp.buildInfo(); buildInfo != nil {
		resp.Data = buildInfo
	}

	json.NewEncoder(w).Encode(resp)
}

func (bp *Blueprint) buildInfo() *BuildInfo {
	info, ok := debug.ReadBuildInfo()

	if !ok {
		return nil
	}

	buildInfo := new(BuildInfo)
	buildInfo.Name = info.Main.Path
	buildInfo.Version = bp.gitRevision(info.Settings)

	return buildInfo
}

func (bp *Blueprint) gitRevision(settings []debug.BuildSetting) string {
	for _, setting := range settings {
		if setting.Key == "vcs.revision" {
			return setting.Value
		}
	}

	return "n/a"
}

// NewBlueprint returns a new initialized Blueprint object.
func NewBlueprint() *Blueprint {
	return new(Blueprint)
}

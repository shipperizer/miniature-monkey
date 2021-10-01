package core

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/zap"

	"github.com/shipperizer/miniature-monkey/config"
	"github.com/shipperizer/miniature-monkey/monitoring"
	"github.com/shipperizer/miniature-monkey/status"
	"github.com/shipperizer/miniature-monkey/webiface"
)

// API is the main object to create a web application
// has the config as attribute and it is basically a wrapper around a mux.Router
// with helpers method to add endpoints (grouped via BlueprintInterface)
// also monitor andlogger are attributes
type API struct {
	name string
	cfg  config.APIConfigInterface

	router *mux.Router

	monitor monitoring.MonitorInterface
	logger  *zap.SugaredLogger
}

// appplyCORS add a CORS wrapper arount the passed in handler if the API has CORS enabled
func (a *API) applyCORS(h http.Handler) http.Handler {
	if cfg := a.cfg.GetCORSConfig(); cfg == nil {
		return h
	} else {
		origins := cfg.GetOrigins()

		c := cors.New(
			cors.Options{
				AllowedOrigins: origins,
				AllowedMethods: []string{
					http.MethodHead,
					http.MethodGet,
					http.MethodPost,
					http.MethodPut,
					http.MethodPatch,
					http.MethodDelete,
				},
				AllowedHeaders:   []string{"*"},
				AllowCredentials: true,
			},
		)
		return c.Handler(h)
	}
}

// Handler returns an http handler created from the main router
func (a *API) Handler() http.Handler {
	return a.applyCORS(a.router)
}

// Router returns the main router
func (a *API) Router() *mux.Router {
	return a.router
}

// UseMiddlewares will apply all the middleware functions to the passed in router
// if no router r is passed, then the main one is used
func (a *API) UseMiddlewares(r *mux.Router, mwf ...mux.MiddlewareFunc) {
	if r == nil {
		r = a.router
	}

	r.Use(mwf...)
}

// UseMiddlewares will register all the blueprints to the passed in router
// if no router r is passed, then the main one is used
func (a *API) RegisterBlueprints(r *mux.Router, blueprints ...webiface.BlueprintInterface) {
	if r == nil {
		r = a.router
	}

	for _, bp := range blueprints {
		bp.Routes(r)
	}
}

// NewAPI returns a new API object implementing webiface.APIInterface
// by default the monitoring and status blueprints are registered and the APITimer and APICount middleware are applied
func NewAPI(cfg config.APIConfigInterface) webiface.APIInterface {
	api := new(API)
	api.name = cfg.GetServiceName()
	api.cfg = cfg
	api.monitor = cfg.GetMonitor()
	api.router = mux.NewRouter()
	api.logger = cfg.GetLogger()

	// register monitoring blueprint by default
	api.RegisterBlueprints(
		nil,
		monitoring.NewBlueprint(),
		status.NewBlueprint(),
	)

	// apply API timer and count middlewares by default
	if api.monitor != nil {
		mdw := monitoring.NewMiddleware(api.monitor, api.monitor.GetService())
		mdws := make([]mux.MiddlewareFunc, 0)
		mdws = append(mdws, mdw.APITime(), mdw.APICount())
		api.UseMiddlewares(nil, mdws...)
	}

	return api
}

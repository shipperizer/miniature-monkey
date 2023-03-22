package core

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	cors "github.com/go-chi/cors"
	trace "go.opentelemetry.io/otel/trace"

	"github.com/shipperizer/miniature-monkey/v2/logging"
	"github.com/shipperizer/miniature-monkey/v2/middlewares/generics"
	core "github.com/shipperizer/miniature-monkey/v2/monitoring/core"
	web "github.com/shipperizer/miniature-monkey/v2/monitoring/web"
	"github.com/shipperizer/miniature-monkey/v2/status"
	"github.com/shipperizer/miniature-monkey/v2/webiface"
)

// API is the main object to create a web application
// has the config as attribute and it is basically a wrapper around a mux.Router
// with helpers method to add endpoints (grouped via BlueprintInterface)
// also monitor andlogger are attributes
type API struct {
	name string
	cfg  APIConfigInterface

	router *chi.Mux

	tracer  trace.Tracer
	monitor core.MonitorInterface
	logger  logging.LoggerInterface
}

func (a *API) middlewareCORS(origins []string) func(http.Handler) http.Handler {
	return cors.Handler(
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
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		},
	)
}

// Handler returns an http handler created from the main router
func (a *API) Handler() http.Handler {
	return a.router
}

// Router returns the main router
func (a *API) Router() *chi.Mux {
	return a.router
}

// UseMiddlewares will apply all the middleware functions to the passed in router
// if no router r is passed, then the main one is used
func (a *API) UseMiddlewares(r *chi.Mux, mwf ...func(http.Handler) http.Handler) {
	if r == nil {
		r = a.router
	}

	r.Use(mwf...)
}

// UseMiddlewares will register all the blueprints to the passed in router
// if no router r is passed, then the main one is used
func (a *API) RegisterBlueprints(r *chi.Mux, blueprints ...webiface.BlueprintInterface) {
	if r == nil {
		r = a.router
	}

	for _, bp := range blueprints {
		bp.Routes(r)
	}
}

// NewAPI returns a new API object implementing webiface.APIInterface
// by default the monitoring and status blueprints are registered and the APITimer and APICount middleware are applied
func NewAPI(cfg APIConfigInterface) *API {
	api := new(API)
	api.name = cfg.GetServiceName()
	api.cfg = cfg
	api.monitor = cfg.GetMonitor()
	api.tracer = cfg.GetTracer()
	api.router = chi.NewMux()
	api.logger = cfg.GetLogger()

	mdws := make(chi.Middlewares, 0)

	// TODO @shipperizer use a better middleware
	// if api.tracer != nil {
	// 	mdws = append(mdws, otelchi.Middleware(api.name, otelchi.WithChiRoutes(api.router)))
	// }

	if api.monitor != nil {
		mdw := generics.NewMiddleware(api.monitor, api.logger)
		mdws = append(mdws, mdw.APITime())
	}

	if c := api.cfg.GetCORSConfig(); c != nil {

		mdws = append(mdws, api.middlewareCORS(c.GetOrigins()))
	}

	api.UseMiddlewares(nil, mdws...)

	// register monitoring blueprint by default
	api.RegisterBlueprints(
		nil,
		web.NewBlueprint(),
		status.NewBlueprint(),
	)

	return api
}

package webiface

import (
	"net/http"

	"github.com/gorilla/mux"
)

// APIInterface
type APIInterface interface {
	Router() *mux.Router
	Handler() http.Handler
	RegisterBlueprints(r *mux.Router, blueprints ...BlueprintInterface)
	UseMiddlewares(r *mux.Router, mwf ...mux.MiddlewareFunc)
}

// BlueprintInterface
type BlueprintInterface interface {
	Routes(router *mux.Router)
}

package webiface

import (
	"net/http"

	"github.com/gorilla/mux"
)

type APIInterface interface {
	Router() *mux.Router
	Handler() http.Handler
	RegisterBlueprints(r *mux.Router, blueprints ...BlueprintInterface)
	UseMiddlewares(r *mux.Router, mwf ...mux.MiddlewareFunc)
}

type BlueprintInterface interface {
	Routes(router *mux.Router)
}

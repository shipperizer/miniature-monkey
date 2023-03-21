package webiface

import (
	chi "github.com/go-chi/chi/v5"
)

// BlueprintInterface
type BlueprintInterface interface {
	Routes(router *chi.Mux)
}

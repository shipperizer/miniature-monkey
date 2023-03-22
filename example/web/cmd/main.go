package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	chi "github.com/go-chi/chi/v5"

	"github.com/kelseyhightower/envconfig"
	"github.com/shipperizer/miniature-monkey/v2/config"
	"github.com/shipperizer/miniature-monkey/v2/core"
	monConfig "github.com/shipperizer/miniature-monkey/v2/monitoring/config"
	monCore "github.com/shipperizer/miniature-monkey/v2/monitoring/core"
	"go.uber.org/zap"
)

type EnvSpec struct {
	Port string `envconfig:"http_port" default:"8000"`
}

type EchoBlueprint struct{}

func (b *EchoBlueprint) Routes(router *chi.Mux) {
	router.Get(
		"/api/v0/echo",
		func(w http.ResponseWriter, r *http.Request) {
			// an example API handler
			json.NewEncoder(w).Encode(map[string]bool{"echo": true})
		},
	)
}

func main() {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()

	if err != nil {
		panic(err.Error())
	}

	var specs EnvSpec
	err = envconfig.Process("", &specs)

	if err != nil {
		logger.Sugar().Fatal(err.Error())
	}

	monitor := monCore.NewMonitor(
		monConfig.NewMonitorConfig("web", nil, logger.Sugar()),
	)

	apiCfg := config.NewAPIConfig(
		"web",
		nil,
		monitor,
		logger.Sugar(),
	)

	api := core.NewAPI(apiCfg)

	api.RegisterBlueprints(api.Router(), new(EchoBlueprint))

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", specs.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      api.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Sugar().Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logger.Sugar().Info("Shutting down")
	os.Exit(0)
}

# Miniature Monkey: small http framework layer built on top of gorilla/mux 

[![test](https://github.com/shipperizer/miniature-monkey/actions/workflows/ci.yaml/badge.svg)](https://github.com/shipperizer/miniature-monkey/actions/workflows/ci.yaml)
[![release](https://github.com/shipperizer/miniature-monkey/actions/workflows/release.yaml/badge.svg)](https://github.com/shipperizer/miniature-monkey/actions/workflows/release.yaml)
[![CodeQL](https://github.com/shipperizer/miniature-monkey/actions/workflows/codeql-analysis.yaml/badge.svg)](https://github.com/shipperizer/miniature-monkey/actions/workflows/codeql-analysis.yaml)
[![codecov](https://codecov.io/gh/shipperizer/miniature-monkey/badge/branch/main/graph/badge.svg)](https://codecov.io/gh/shipperizer/miniature-monkey/badge)
[![Go Reference](https://pkg.go.dev/badge/github.com/shipperizer/miniature-monkey.svg)](https://pkg.go.dev/github.com/shipperizer/miniature-monkey)

Miniature Monkey simply bootstraps a simple http application with `gorilla/mux`, adding a couple of endpoints like status and the prometheus metrics handler.

The `API` struct offers a set of methods to add endpoints via the `BlueprintInterface` abstraction, also you can register Middleware functions 


## Examples

```
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shipperizer/miniature-monkey/config"
	"github.com/shipperizer/miniature-monkey/core"
	"github.com/shipperizer/miniature-monkey/utils"
	monitoringLibrary "github.com/some/monitoring/library"
)

func main() {
	monitor := monitoringLibrary.NewMonitor()
	logger := utils.NewLogger("info")

	apiCfg := config.NewAPIConfig(
		"test-api",
		config.NewCORSConfig("google.com"),
		monitor,
		logger,
	)

	srv := &http.Server{
		Addr: "0.0.0.0:8000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      core.NewApi(apiCfg).Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

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
	logger.Info("Shutting down")
	os.Exit(0)
}
```
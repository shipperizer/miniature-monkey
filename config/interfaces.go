package config

import (
	"github.com/shipperizer/miniature-monkey/monitoring"
	"go.uber.org/zap"
)

type CORSConfigInterface interface {
	GetOrigins() []string
}

type APIConfigInterface interface {
	GetServiceName() string
	GetMonitor() monitoring.MonitorInterface
	GetCORSConfig() CORSConfigInterface
	GetLogger() *zap.SugaredLogger
}

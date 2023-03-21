// Interface definitions
package core

import (
	"github.com/shipperizer/miniature-monkey/v2/config"
	"github.com/shipperizer/miniature-monkey/v2/logging"
	core "github.com/shipperizer/miniature-monkey/v2/monitoring/core"
)

type APIConfigInterface interface {
	GetServiceName() string
	GetMonitor() core.MonitorInterface
	GetCORSConfig() config.CORSConfigInterface
	GetLogger() logging.LoggerInterface
}

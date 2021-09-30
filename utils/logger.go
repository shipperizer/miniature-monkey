package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

// NewLogger creates a new default logger
// it will need to be closed with
// ```
// defer logger.Desugar().Sync()
// ```
// to make sure all has been piped out before terminating
func NewLogger(l string) *zap.SugaredLogger {
	var lvl string

	switch val := strings.ToLower(l); {
	case val == "debug":
	case val == "error":
	case val == "warning":
	case val == "info":
		lvl = val
	default:
		lvl = "info"
	}

	rawJSON := []byte(
		fmt.Sprintf(
			`{
				"level": "%s",
				"encoding": "json",
				"outputPaths": ["stdout"],
				"errorOutputPaths": ["stdout","stderr"],
				"encoderConfig": {
				"messageKey": "message",
				"levelKey": "severity",
				"levelEncoder": "lowercase",
				"timeKey": "@timestamp",
				"timeEncoder": "rfc3339nano"
				}
			}`,
			lvl),
	)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger.Sugar()
}

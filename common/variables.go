package common

import (
	log "github.com/sirupsen/logrus"

	"github.com/asdine/storm/v3"
	"github.com/shelly-tools/core/config"
)

var (
	Config      *config.Config
	LogInstance *log.Logger
	DB          *storm.DB
)

func PrepareLogInstance() {
	LogInstance = log.New()

	LogInstance.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	LogInstance.SetLevel(log.ErrorLevel)
}

func ChangeLogLevel(level string) {
	var logLevel log.Level
	switch level {
	case "debug":
		logLevel = log.DebugLevel
	case "info":
		logLevel = log.InfoLevel
	case "error":
		logLevel = log.ErrorLevel
	default:
		logLevel = log.DebugLevel
	}
	LogInstance.SetLevel(logLevel)
}

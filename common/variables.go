package common

import (
	log "github.com/sirupsen/logrus"

	"github.com/shelly-tools/core/config"
)

var (
	Config      *config.Config
	LogInstance *log.Logger
)

package logs

import (
	"github.com/daqiancode/env"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	level, err := zerolog.ParseLevel(env.Get("LOG_LEVEL", "DEBUG"))
	if err != nil {
		panic(err)
	}
	zerolog.SetGlobalLevel(level)
}

func GetLogger(file string) zerolog.Logger {
	return log.Logger.With().Str("file", file).Logger()
}

var Log = log.Logger

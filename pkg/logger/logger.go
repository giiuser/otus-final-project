package logger

import (
	"errors"
	"github.com/rs/zerolog"
	"os"

	"github.com/giiuser/otus-final-project/pkg/config"
	"github.com/rs/zerolog/log"
)

var ErrFileLog = errors.New("cannot setup file log")

func getLogLevel(str string) zerolog.Level {
	switch str {
	case "error":
		return zerolog.ErrorLevel
	case "warn":
		return zerolog.WarnLevel
	case "info":
		return zerolog.InfoLevel
	case "debug":
		return zerolog.DebugLevel
	default:
		return zerolog.InfoLevel
	}
}

func Init(c *config.Config) {
	zerolog.SetGlobalLevel(getLogLevel(c.LogLevel))
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: zerolog.TimeFieldFormat})
}

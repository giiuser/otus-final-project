package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"

	"github.com/giiuser/otus-final-project/internal/app"
	"github.com/giiuser/otus-final-project/internal/config"
	"github.com/giiuser/otus-final-project/internal/logger"
	"github.com/rs/zerolog/log"
)

var cfgPath string

var ErrAppFatal = errors.New("application cannot start")

func init() {
	flag.StringVar(&cfgPath, "config", "", "Image previewer config")
}

func main() {
	flag.Parse()

	cfg, err := config.Read(cfgPath)
	if err != nil {
		log.Fatal().Err(err).Msgf("%s", ErrAppFatal)
	}
	logger.Init(cfg)
	log.Debug().Msgf("Config Init %+v", cfg)
	app, err := app.New(cfg, http.DefaultClient)
	if err != nil {
		log.Fatal().Err(err).Msgf("%s", ErrAppFatal)
	}
	err = app.Run(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal().Err(err).Msgf("%s", ErrAppFatal)
	}
}

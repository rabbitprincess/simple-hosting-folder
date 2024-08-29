package app

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

func NewApp(logger zerolog.Logger, mode string) *App {
	return &App{
		logger: logger,
		mode:   mode,
	}
}

type App struct {
	server *http.Server
	logger zerolog.Logger
	mode   string
}

func (a *App) Run(dataPath, certPath, keyPath string) error {
	fs := http.FileServer(http.Dir(dataPath))
	http.Handle("/", fs)

	a.server = &http.Server{
		Addr:         ":443",
		Handler:      http.DefaultServeMux,
		TLSConfig:    &tls.Config{MinVersion: tls.VersionTLS12},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	a.logger.Info().Msg("Starting server on :443")

	if err := a.server.ListenAndServeTLS(certPath, keyPath); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("could not listen on %s: %w", a.server.Addr, err)
	}

	return nil
}

func (a *App) Stop() error {
	a.logger.Info().Msg("Server stopped gracefully...")
	return nil
}

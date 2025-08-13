// SPDX-License-Identifier: EUPL-1.2

package api

import (
	"azugo.io/azugo"
	"azugo.io/azugo/server"
	"azugo.io/core/instrumenter"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nobid-lsp-latvia/go-audit"
	"github.com/nobid-lsp-latvia/lx-go-jsondb"
	"github.com/spf13/cobra"
)

// App is the application instance.
type App struct {
	*azugo.App

	config *Configuration
	store  jsondb.Store
	db     *pgxpool.Pool
	audit  audit.Audit
}

// New returns a new application instance.
func New(cmd *cobra.Command, version string) (*App, error) {
	config := NewConfiguration()

	a, err := server.New(cmd, server.Options{
		AppName:       "Admin API",
		AppVer:        version,
		Configuration: config,
	})
	if err != nil {
		return nil, err
	}

	store, db, err := jsondb.New(a.App, config.Postgres)
	if err != nil {
		return nil, err
	}

	app := &App{
		App:    a,
		config: config,
		store:  store,
		db:     db,
		audit:  audit.New(config.Audit, nil),
	}

	app.Instrumentation(instrumenter.CombinedInstrumenter(app.Instrumenter(), app.storeInstrumenter()))

	return app, nil
}

// Start starts the application.
func (a *App) Start() error {
	if err := a.Store().Start(a.BackgroundContext()); err != nil {
		return err
	}

	return a.App.Start()
}

// Config returns application configuration.
//
// Panics if configuration is not loaded.
func (a *App) Config() *Configuration {
	if a.config == nil || !a.config.Ready() {
		panic("configuration is not loaded")
	}

	return a.config
}

func (a *App) Store() jsondb.Store {
	return a.store
}

// Audit returns audit service interface.
func (a *App) Audit() audit.Audit {
	return a.audit
}

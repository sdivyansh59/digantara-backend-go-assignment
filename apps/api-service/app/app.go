package app

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/setup"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/setup/dbconfig"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/internal-lib/utils"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/routes"
	"github.com/uptrace/bun"
)

// App is the main application struct
type App struct {
	*utils.WithLogger
	router      *chi.Mux
	huma        *huma.API
	schedulerDB *bun.DB
	controllers *setup.Controllers
	config      *utils.DefaultConfig
}

func newApp(r *chi.Mux, h *huma.API, config *utils.DefaultConfig, c *setup.Controllers, logger *utils.WithLogger, jobSchedulerDB *dbconfig.JobSchedulerDB) *App {
	return &App{
		WithLogger:  logger,
		router:      r,
		huma:        h,
		schedulerDB: jobSchedulerDB.DB,
		controllers: c,
		config:      config,
	}
}

// Run starts the application server
func (a *App) Run() error {
	// Configure your routes
	a.registerRoutes()

	// Start the HTTP server
	log.Info().Msgf("Starting server on %s", a.config.HTTPAddress)
	return http.ListenAndServe(a.config.HTTPAddress, a.router)
}

// registerRoutes configures all API endpoints
func (a *App) registerRoutes() {
	if a.huma == nil {
		log.Fatal().Msgf("huma is nil")
	}

	routes.RegisterRoutes(a.huma, a.controllers)
}

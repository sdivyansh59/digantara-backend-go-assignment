//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/job"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/scheduler"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/setup"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/setup/dbconfig"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/internal-lib/utils"
)

// InitializeApp wires up all dependencies and returns the application/service instance
func InitializeApp() (*App, error) {
	wire.Build(
		// Core configuration - must be first
		utils.ProvideDefaultConfig,

		// Initialize global logger early (depends on config, returns logger instance)
		utils.InitGlobalLogger,
		utils.NewWithLogger,

		// Database initialization with migrations
		dbconfig.ProvideJobSchedulerDB,

		// Infrastructure
		setup.ProvideSingletonChiRouter,
		setup.ProvideSingletonHuma,
		setup.ProvideSnowflakeGenerator,
		setup.ProvideControllers,

		// Application
		newApp,

		// Initialize application controllers, converter and repositories
		// job
		job.NewController,
		job.NewConverter,
		job.NewRepository,
		// scheduler
		scheduler.NewController,
	)
	return nil, nil
}

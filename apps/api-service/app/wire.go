//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/greeting"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/internal-lib/snowflake"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/internal-lib/utils"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/setup"
)

// InitializeApp wires up all dependencies and returns the application/service instance
func InitializeApp() (*App, error) {
	wire.Build(
		setup.ProvideSingletonChiRouter,
		setup.ProvideSingletonHuma,
		utils.ProvideDefaultConfig,
		setup.ProvideControllers,
		snowflake.NewGenerator(utils.GetEnvOr("MACHINE_ID", 1)),
		newApp,

		// initialize all domains controller and repository
		greeting.NewController,
		//greeting.NewRepository,
	)
	return nil, nil
}

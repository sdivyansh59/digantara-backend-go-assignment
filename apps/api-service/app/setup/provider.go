package setup

import (
	"sync"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/greeting"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/internal-lib/snowflake"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/internal-lib/utils"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/job"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/scheduler"
)

var (
	humaOnce     sync.Once
	humaInstance *huma.API
)

// ProvideSingletonHuma returns a singleton Huma API instance
func ProvideSingletonHuma(router *chi.Mux) *huma.API {
	humaOnce.Do(func() {
		api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))
		humaInstance = utils.ToPointer(api)
	})
	return humaInstance
}

// Controllers holds all application controllers
type Controllers struct {
	Greeting  *greeting.Controller
	Job       *job.Controller
	Scheduler *scheduler.Controller
	// Add other controllers here as you build them
	// User     *user.Controller
	// Product  *product.Controller
}

// ProvideControllers wires up all controllers
func ProvideControllers(
	greetingController *greeting.Controller,
	jobController *job.Controller,
	// Add other controllers here as parameters
) *Controllers {
	return &Controllers{
		Greeting: greetingController,
		Job:      jobController,
		// Add other controllers
	}
}

// ProvideSnowflakeGenerator provides a snowflake ID generator
func ProvideSnowflakeGenerator() (*snowflake.Generator, error) {
	machineID := utils.GetEnvOrInt64("MACHINE_ID", 1)
	return snowflake.NewGenerator(machineID)
}

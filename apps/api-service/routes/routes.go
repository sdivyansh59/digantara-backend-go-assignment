package routes

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/setup"
)

// RegisterRoutes registers all job routes to the API
func RegisterRoutes(api *huma.API, c *setup.Controllers) {
	huma.Register(*api, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
		Summary:     "Get a greeting",
		Description: "Get a greeting for a person by name.",
		Tags:        []string{"Greetings"},
	}, c.Greeting.GetGreeting)

	huma.Register(*api, huma.Operation{
		OperationID:   "post-review",
		Method:        http.MethodPost,
		Path:          "/reviews",
		Summary:       "Post a review",
		Tags:          []string{"Reviews"},
		DefaultStatus: http.StatusCreated,
	}, c.Greeting.PostReview)

	// Job routes
	huma.Register(*api, huma.Operation{
		OperationID: "create-job",
		Method:      http.MethodPost,
		Path:        "/jobs",
		Summary:     "Create a new job",
		Description: "Create a new job with a name, optional description, scheduled time, and creator email. " +
			"The scheduled time must be a future Unix timestamp.",
		Tags:          []string{"Jobs"},
		DefaultStatus: http.StatusCreated,
	}, c.Job.CreateJob)

	huma.Register(*api, huma.Operation{
		OperationID: "get-all-jobs",
		Method:      http.MethodGet,
		Path:        "/jobs",
		Summary:     "Get all jobs",
		Description: "Retrieve a list of all jobs.",
		Tags:        []string{"Jobs"},
	}, c.Job.FilterJobs)

	huma.Register(*api, huma.Operation{
		OperationID: "get-job-by-id",
		Method:      http.MethodGet,
		Path:        "/jobs/{id}",
		Summary:     "Get job by ID",
		Description: "Retrieve a job by its unique identifier.",
		Tags:        []string{"Jobs"},
	}, c.Job.GetJobByID)

	huma.Register(*api, huma.Operation{
		OperationID: "delete-job-by-id",
		Method:      http.MethodDelete,
		Path:        "/jobs/{id}",
		Summary:     "Get job by ID",
		Description: "Retrieve a job by its unique identifier.",
		Tags:        []string{"Jobs"},
	}, c.Job.DeleteJobByID)

}

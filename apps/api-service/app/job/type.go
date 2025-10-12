package job

import (
	"time"

	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/shared"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/internal-lib/snowflake"
)

type Job struct {
	id          snowflake.ID
	name        string
	description *string
	status      shared.JobStatus
	interval    bool  // default false
	scheduledAt int64 // Unix timestamp, 8 bytes, efficient for sorting and db indexing
	lastRunAt   *time.Time
	createdBy   string
	createdAt   time.Time
	updatedAt   time.Time
}

type GetJobByIDInput struct {
	ID string `path:"id" validate:"required,uuid" doc:"Unique identifier of the job"`
}

type CreateJobInput struct {
	Name        string  `json:"name" validate:"required,min=3,max=100" doc:"Job name"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=500" doc:"Job description"`
	Interval    bool    `json:"interval" validate:"-" doc:"Indicates if the job is recurring (default: false)" example:"false"`
	ScheduledAt int64   `json:"scheduled_at" validate:"required,gt=0" doc:"Scheduled time of the Job (Unix timestamp, must be in the future)" example:"1728691200"` // Unix timestamp
	CreatedBy   string  `json:"created_by" validate:"required,email" doc:"Email of the job creator"`
}

type FilterJobsInput struct{}

type DeleteJobByIDInput struct {
	ID string `path:"id" validate:"required,uuid" doc:"Unique identifier of the job to delete"`
}

type DeleteJobByIDResponse struct {
	Success bool `json:"success" doc:"Indicates if the job was successfully deleted"`
}

type FilterJobsResponse struct {
	Jobs []JobResponse `json:"jobs" doc:"List of jobs"`
}

type JobResponse struct {
	ID          string           `json:"id" doc:"Unique identifier of the created job"`
	Name        string           `json:"name" doc:"Name of the created job"`
	Description *string          `json:"description,omitempty" doc:"Description of the created job"`
	Status      shared.JobStatus `json:"status" doc:"Current status of the job" enum:"PENDING,RUNNING,COMPLETED,FAILED"`
	Interval    bool             `json:"interval" doc:"Indicates if the job is recurring"`
	ScheduledAt int64            `json:"scheduled_at" doc:"Scheduled time of the job (Unix timestamp)"`
	LastRunAt   *time.Time       `json:"last_run_at,omitempty" doc:"Last run time of the job"`
	CreatedBy   string           `json:"created_by" doc:"Email of the job creator"`
	CreatedAt   time.Time        `json:"created_at" doc:"Creation time of the job (Unix timestamp)"`
	UpdatedAt   time.Time        `json:"updated_at" doc:"Last update time of the job (Unix timestamp)"`
}

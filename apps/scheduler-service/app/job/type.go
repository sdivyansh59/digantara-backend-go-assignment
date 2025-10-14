package job

import (
	"time"

	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/shared"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/internal-lib/snowflake"
	"github.com/uptrace/bun"
)

type Job struct {
	bun.BaseModel `bun:"table:job,alias:job"`

	Id             snowflake.ID           `bun:"id,pk,notnull"`
	Name           string                 `bun:"name,notnull"`
	Description    *string                `bun:"description"`
	Status         shared.JobStatus       `bun:"status,notnull"`
	IntervalTime   *int64                 `bun:"interval_time"`        // nullable for one-time jobs
	ScheduledAt    int64                  `bun:"scheduled_at,notnull"` // Unix timestamp in milliseconds
	LastRunAt      *time.Time             `bun:"last_run_at"`
	SuccessfulRuns int                    `bun:"successful_runs,notnull,default:0"`
	Attributes     map[string]interface{} `bun:"attributes,type:jsonb"` // explicitly specify JSONB type
	CreatedBy      string                 `bun:"created_by,notnull"`
	CreatedAt      time.Time              `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt      time.Time              `bun:"updated_at,notnull,default:current_timestamp"`
}

type GetJobByIDInput struct {
	ID string `path:"id" validate:"required,uuid" doc:"Unique identifier of the job"`
}

type CreateJobInput struct {
	Name         string                 `json:"name" validate:"required,min=3,max=100" doc:"Job name"`
	Description  *string                `json:"description,omitempty" validate:"omitempty,max=500" doc:"Job description"`
	Interval     bool                   `json:"interval" validate:"-" doc:"Indicates if the job is recurring (default: false)" example:"false"`
	IntervalTime *int64                 `json:"interval_time,omitempty" doc:"Interval time in minutes (for recurring jobs)" example:"1440 mins for a day"`
	ScheduledAt  int64                  `json:"scheduled_at" validate:"required,gt=0" doc:"Scheduled time of the Job (Unix timestamp, must be in the future)" example:"1728691200"` // Unix timestamp
	Attributes   map[string]interface{} `json:"attributes,omitempty" validate:"-" doc:"Custom job attributes (flexible key-value pairs)" example:"{\"priority\":\"high\",\"department\":\"engineering\",\"tags\":[\"critical\",\"backend\"]}"`
	CreatedBy    string                 `json:"created_by" validate:"required,email" doc:"Email of the job creator"`
}

type FilterJobsInput struct{}

type DeleteJobByIDInput struct {
	ID string `path:"id" validate:"required,uuid" doc:"Unique identifier of the job to delete"`
}

type JobDTO struct {
	ID             string                 `json:"id" doc:"Unique identifier of the created job"`
	Name           string                 `json:"name" doc:"Name of the created job"`
	Description    *string                `json:"description,omitempty" doc:"Description of the created job"`
	Status         shared.JobStatus       `json:"job_status" doc:"Current status of the job" enum:"PENDING,RUNNING,COMPLETED,FAILED"`
	IntervalTime   int64                  `json:"interval_time" doc:"Interval time in minutes (for recurring jobs)" example:"1440 mins for a day"`
	ScheduledAt    int64                  `json:"scheduled_at" doc:"Scheduled time of the job (Unix timestamp)"`
	LastRunAt      *time.Time             `json:"last_run_at,omitempty" doc:"Last run time of the job"`
	Attributes     map[string]interface{} `json:"attributes,omitempty" doc:"Custom job attributes"`
	SuccessfulRuns int                    `json:"successful_runs" doc:"Number of successful runs for the job"`
	CreatedBy      string                 `json:"created_by" doc:"Email of the job creator"`
	CreatedAt      time.Time              `json:"created_at" doc:"Creation time of the job (Unix timestamp)"`
	UpdatedAt      time.Time              `json:"updated_at" doc:"Last update time of the job (Unix timestamp)"`
}

// Huma response wrappers

type CreateJobResponse struct {
	Body JobDTO
}

type GetJobByIDResponse struct {
	Body JobDTO
}

type FilterJobsResponse struct {
	Body struct {
		Jobs []JobDTO `json:"jobs" doc:"List of jobs"`
	}
}

type DeleteJobResponse struct {
	Body struct {
		Success bool `json:"success" doc:"Indicates if the job was successfully deleted"`
	}
}

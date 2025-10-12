package job

import (
	"context"
	"fmt"
	"time"

	"github.com/sdivyansh59/digantara-backend-golang-assignment/internal-lib/snowflake"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/internal-lib/utils"
)

type Controller struct {
	*utils.WithLogger
	snowflake  *snowflake.Generator
	converter  *Converter
	repository IRepository
}

func NewController(logger *utils.WithLogger, snowflake *snowflake.Generator, converter *Converter,
	repository IRepository) *Controller {
	return &Controller{
		WithLogger: logger,
		snowflake:  snowflake,
		converter:  converter,
		repository: repository,
	}
}

func (c *Controller) FilterJobs(ctx context.Context, input *FilterJobsInput) (*FilterJobsResponse, error) {
	if !c.isAuthorized(ctx) {
		return nil, fmt.Errorf("unauthorized: you do not have permission to delete this job")
	}

	entities, err := c.repository.Filter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to filter jobs: %w", err)
	}

	var jobs []JobResponse
	for _, entity := range entities {
		jobs = append(jobs, *c.converter.ToDTO(&entity))
	}

	return &FilterJobsResponse{
		Jobs: jobs,
	}, nil
}

func (c *Controller) GetJobByID(ctx context.Context, input *GetJobByIDInput) (*JobResponse, error) {
	if !c.isAuthorized(ctx) {
		return nil, fmt.Errorf("unauthorized: you do not have permission to delete this job")
	}

	// validate job's id
	jobID, err := snowflake.ConvertToSnowflake(input.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid job ID: %w", err)
	}

	// Check if job exists
	job, err := c.repository.GetByID(ctx, jobID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve job: %w", err)
	}
	if job == nil {
		return nil, fmt.Errorf("job not found")
	}

	return c.converter.ToDTO(job), nil
}

func (c *Controller) CreateJob(ctx context.Context, input *CreateJobInput) (*JobResponse, error) {
	if !c.isAuthorized(ctx) {
		return nil, fmt.Errorf("unauthorized: you do not have permission to create a job")
	}

	// Validate that scheduled time is in the future
	currentTime := time.Now().Unix()
	if input.ScheduledAt <= currentTime {
		return nil, fmt.Errorf("scheduled_at must be a future timestamp (current: %d, provided: %d)", currentTime, input.ScheduledAt)
	}

	entity := c.converter.ToEntity(input)
	err := c.repository.Create(ctx, entity)
	if err != nil {
		return nil, fmt.Errorf("failed to create job: %w", err)
	}

	return c.converter.ToDTO(entity), nil
}

func (c *Controller) DeleteJobByID(ctx context.Context, input *DeleteJobByIDInput) (*DeleteJobByIDResponse, error) {
	if !c.isAuthorized(ctx) {
		return nil, fmt.Errorf("unauthorized: you do not have permission to delete this job")
	}

	// validate job's id
	jobID, err := snowflake.ConvertToSnowflake(input.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid job ID: %w", err)
	}

	// Check if job with id exists
	job, err := c.repository.GetByID(ctx, jobID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve job: %w", err)
	}
	if job == nil {
		return nil, fmt.Errorf("job not found")
	}

	// Proceed to delete the job
	err = c.repository.DeleteByID(ctx, job)
	if err != nil {
		return nil, fmt.Errorf("failed to delete job: %w", err)
	}

	// Return success response
	return &DeleteJobByIDResponse{
		Success: true,
	}, nil
}

package job

import (
	"context"
	"fmt"
	"time"

	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/shared"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/internal-lib/snowflake"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/internal-lib/utils"
)

type Controller struct {
	*utils.WithLogger
	snowflake  *snowflake.Generator
	converter  *Converter
	repository IRepository
	wakeupChan chan *shared.WakeupEvent
}

func NewController(logger *utils.WithLogger, snowflake *snowflake.Generator, converter *Converter,
	repository IRepository, wakeupChan chan *shared.WakeupEvent) *Controller {
	return &Controller{
		WithLogger: logger,
		snowflake:  snowflake,
		converter:  converter,
		repository: repository,
		wakeupChan: wakeupChan,
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

	jobs := make([]JobDTO, 0, len(entities))
	for _, entity := range entities {
		jobs = append(jobs, *c.converter.ToDTO(&entity))
	}

	resp := &FilterJobsResponse{}
	resp.Body.Jobs = jobs
	return resp, nil
}

func (c *Controller) GetJobByID(ctx context.Context, input *GetJobByIDInput) (*GetJobByIDResponse, error) {
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

	return &GetJobByIDResponse{
		Body: *c.converter.ToDTO(job),
	}, nil
}

func (c *Controller) CreateJob(ctx context.Context, input *CreateJobInput) (*CreateJobResponse, error) {
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

	// Notify scheduler of new job (non-blocking)
	select {
	case c.wakeupChan <- &shared.WakeupEvent{
		JobID:       entity.Id,
		ScheduledAt: entity.ScheduledAt,
	}:
		c.Logger.Info().Msgf("Notified scheduler of new job %s", entity.Id)
	default:
		c.Logger.Warn().Msg("Scheduler wakeup channel is full, skipping notification")
	}

	return &CreateJobResponse{
		Body: *c.converter.ToDTO(entity),
	}, nil
}

func (c *Controller) DeleteJobByID(ctx context.Context, input *DeleteJobByIDInput) (*DeleteJobResponse, error) {
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

	resp := &DeleteJobResponse{}
	resp.Body.Success = true
	return resp, nil
}

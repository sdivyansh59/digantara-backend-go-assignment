package scheduler

import (
	"context"
	"time"

	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/job"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/shared"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/internal-lib/snowflake"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/internal-lib/utils"
)

const defaultSleepTime = 5 * time.Minute

type Controller struct {
	*utils.WithLogger
	snowflake     *snowflake.Generator
	jobRepository job.IRepository
	jobConverter  *job.Converter
	sleepTime     time.Duration
	wakeupChan    chan *shared.WakeupEvent // Read-only channel
}

func NewController(logger *utils.WithLogger, snowflake *snowflake.Generator, repo job.IRepository,
	converter *job.Converter, wakeupChan chan *shared.WakeupEvent) *Controller {
	return &Controller{
		WithLogger:    logger,
		snowflake:     snowflake,
		jobRepository: repo,
		jobConverter:  converter,
		sleepTime:     1 * time.Minute, // default
		wakeupChan:    wakeupChan,
	}
}

func (c *Controller) findAndUpdateSleepTime(ctx context.Context) {
	scheduledTime, err := c.jobRepository.GetNextJobScheduledTime(ctx)
	if err != nil {
		c.Logger.Error().Err(err)
		c.sleepTime = defaultSleepTime
		return
	}

	if scheduledTime == nil { // may be no jobs scheduled
		c.sleepTime = defaultSleepTime
		return
	}

	currentTime := time.Now().UnixMilli()
	if *scheduledTime <= currentTime {
		c.sleepTime = 0
		return
	}

	sleepDuration := time.Duration(*scheduledTime-currentTime) * time.Millisecond
	c.sleepTime = sleepDuration
}

func (c *Controller) runJob(ctx context.Context, job *job.Job) {
	if job == nil {
		c.Logger.Info().Msg("No job to run at this time")
		return
	}

	// Assuming job will take approx 10sec to execute
	time.Sleep(10 * time.Second)

	// Note: if it failed then go and marked status = FAILED and handle this case separately.

	// If completed
	job.Status = shared.JobStatusCompleted
	job.LastRunAt = utils.ToPointer(time.Now())
	if job.IntervalTime != nil {
		nextScheduledTimeInMins := job.IntervalTime
		// schedule job again for next interval
		job.ScheduledAt = time.Now().Add(time.Duration(*nextScheduledTimeInMins) * time.Minute).UnixMilli()
	}

	err := c.jobRepository.Update(ctx, job)
	if err != nil {
		c.Logger.Error().Err(err).Msgf("error while updating job status to COMPLETED for job id:%s", job.Id)
		// update is as a failed job
		job.Status = shared.JobStatusFailed
		err = c.jobRepository.Update(ctx, job)
		if err != nil {
			c.Logger.Error().Err(err).Msgf("error while updating job status to FAILED for job id:%s", job.Id)
		}

		return
	}

	c.Logger.Info().Msgf("Job with id:%s completed successfully", job.Id)
}

// Scheduler responsible for running scheduled jobs at their scheduled time.
func (c *Controller) Scheduler(ctx context.Context) error {
	go func() {
		for {
			c.findAndUpdateSleepTime(ctx)
			c.Logger.Info().Msgf("Scheduler sleeping for %v", c.sleepTime)

			timer := time.NewTimer(c.sleepTime)

			select {
			case <-timer.C:
				// Normal wakeup after sleep
			case event := <-c.wakeupChan:
				// Early wakeup triggered by new job
				c.Logger.Info().Msgf("Scheduler woken up by job %s scheduled at %d", event.JobID, event.ScheduledAt)
				timer.Stop()

				continue // re-evaluate sleep time immediately
			}

			jobToRun, err := c.jobRepository.GetNextJobToRun(ctx)
			if err != nil {
				c.Logger.Error().Err(err).Msg("Failed to get next job to run")
			}

			go c.runJob(ctx, jobToRun)
		}
	}()

	c.Logger.Info().Msg("Scheduler started")
	return nil
}

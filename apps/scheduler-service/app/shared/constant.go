package shared

import "github.com/sdivyansh59/digantara-backend-golang-assignment/internal-lib/snowflake"

// JobStatus represents the possible states of a job
type JobStatus string

const (
	JobStatusScheduled JobStatus = "SCHEDULED"
	JobStatusRunning   JobStatus = "RUNNING"
	JobStatusCompleted JobStatus = "COMPLETED"
	JobStatusFailed    JobStatus = "FAILED"
)

// WakeupEvent represents an event to wake up the scheduler
type WakeupEvent struct {
	JobID       snowflake.ID
	ScheduledAt int64
}

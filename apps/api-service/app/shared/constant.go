package shared

// JobStatus represents the possible states of a job
type JobStatus string

const (
	JobStatusScheduled JobStatus = "SCHEDULED"
	JobStatusRunning   JobStatus = "RUNNING"
	JobStatusCompleted JobStatus = "COMPLETED"
	JobStatusFailed    JobStatus = "FAILED"
)

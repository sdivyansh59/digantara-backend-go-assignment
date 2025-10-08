package job

import (
	"time"

	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/internal-lib/snowflake"
)

type Job struct {
	id          snowflake.ID
	name        string
	description *string
	status      string

	scheduledAt int64 // Unix timestamp, 8 bytes, efficient for sorting and db indexing
	createdBy   string
	createdAt   time.Time
	updatedAt   time.Time
}

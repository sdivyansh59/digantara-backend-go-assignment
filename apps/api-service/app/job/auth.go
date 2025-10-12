package job

import "context"

func (c *Controller) isAuthorized(ctx context.Context) bool {
	// We saved user info in context in auth middleware
	// ex: userEmail := ctx.Value("userEmail")
	// Here we can check if user has permission to perform the action

	// For simplicity of the assignment, let's assume all authenticated users are authorized
	return true
}

//// Setting scheduledAt
//job.scheduledAt = time.Now().Add(24 * time.Hour).Unix() // Schedule for tomorrow
//
//// Converting back to time.Time when needed
//scheduledTime := time.Unix(job.scheduledAt, 0)
//
//// Easy sorting
//sort.Slice(jobs, func(i, j int) bool {
//	return jobs[i].scheduledAt < jobs[j].scheduledAt
//})

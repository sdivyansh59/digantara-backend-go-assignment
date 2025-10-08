package job

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

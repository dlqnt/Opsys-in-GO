package stride

import (
	"dat320/lab4/scheduler/job"
	"time"
)

// NewJob creates a job for stride scheduling.
func NewJob(size, tickets int, estimated time.Duration) *job.Job {
	const numerator = 10_000
	job_value := job.New(size, estimated)
	if tickets == 0 {
		tickets = 1
	}
	job_value.Stride = numerator / tickets
	job_value.Tickets = tickets
	job_value.Pass = 0
	return job_value
}

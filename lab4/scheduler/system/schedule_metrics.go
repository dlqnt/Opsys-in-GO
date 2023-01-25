package system

import (
	"dat320/lab4/scheduler/job"
	"math"
	"time"
)

// Avg returns the average of a metric defined by the function f.
func (sch Schedule) Avg(f func(*job.Job) time.Duration) time.Duration {
	sum := time.Duration(0)

	for _, value := range sch {
		sum += f(value.Job)
	}
	math.Round(float64(sum))
	sum = sum / time.Duration(len(sch))

	return sum
}

func (sch Schedule) AvgResponseTime() time.Duration {

	sum := time.Duration(0)

	for i := range sch {
		sum += sch[i].Job.ResponseTime()

	}
	math.Round(float64(sum))

	sum = sum / time.Duration(len(sch))
	return sum
}

func (sch Schedule) AvgTurnaroundTime() time.Duration {

	sum := time.Duration(0)

	for i := range sch {
		sum += sch[i].Job.TurnaroundTime()

	}
	math.Round(float64(sum))

	sum = sum / time.Duration(len(sch))
	return sum
}

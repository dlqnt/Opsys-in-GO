package job

import (
	"dat320/lab4/scheduler/system/systime"
	"time"
)

func (j *Job) Scheduled(s systime.SystemTime) {
	j.SystemTime = s
	j.arrival = s.Now()

}

func (j *Job) Started(cpuID int) {
	if j.start == NotStartedYet {
		j.start = j.SystemTime.Now()
	}
}

func (j Job) TurnaroundTime() time.Duration {
	Turn := j.finished - j.arrival
	return Turn
}

func (j Job) ResponseTime() time.Duration {
	Resp := j.start - j.arrival
	return Resp
}

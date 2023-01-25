package rr

import (
	"dat320/lab4/scheduler/cpu"
	"dat320/lab4/scheduler/job"
	"time"
)

type roundRobin struct {
	queue   job.Jobs
	cpu     *cpu.CPU
	quantum time.Duration
}

func New(cpus []*cpu.CPU, quantum time.Duration) *roundRobin {
	if len(cpus) != 1 {
		panic("rr scheduler supports only a single CPU")
	}
	return &roundRobin{
		cpu:     cpus[0],
		queue:   make(job.Jobs, 0),
		quantum: quantum,
	}
}

func (rr *roundRobin) Add(job *job.Job) {
	rr.queue = append(rr.queue, job)
}

// Tick runs the scheduled jobs for the system time, and returns
// the number of jobs finished in this tick. Depending on scheduler requirements,
// the Tick method may assign new jobs to the CPU before returning.
func (rr *roundRobin) Tick(systemTime time.Duration) int {
	jobsFinished := 0
	if rr.cpu.IsRunning() {
		if rr.cpu.Tick() {
			jobsFinished++
			if systemTime%rr.quantum == 0 {
				rr.reassign()
			}
		} else {
			if systemTime%rr.quantum == 0 {
				rr.queue = append(rr.queue, rr.cpu.CurrentJob())
				rr.reassign()
			}
		}
	} else {
		if systemTime%rr.quantum == 0 { // CPU is idle, find new job in own queue
			rr.reassign()
		}
	}
	return jobsFinished
}

// reassign assigns a job to the cpu
func (rr *roundRobin) reassign() {
	nextJob := rr.getNewJob()
	rr.cpu.Assign((nextJob))
}

// getNewJob finds a new job to run on the CPU, removes the job from the queue and returns the job
func (rr *roundRobin) getNewJob() *job.Job {
	if rr.queue.Len() == 0 {
		return nil
	}
	removedJob := rr.queue[0]
	rr.queue = rr.queue[1:]
	return removedJob
}

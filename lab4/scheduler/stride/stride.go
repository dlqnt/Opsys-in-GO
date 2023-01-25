package stride

import (
	"dat320/lab4/scheduler/cpu"
	"dat320/lab4/scheduler/job"
	"time"
)

type stride struct {
	queue   job.Jobs
	cpu     *cpu.CPU
	quantum time.Duration
}

func New(cpus []*cpu.CPU, quantum time.Duration) *stride {
	if len(cpus) != 1 {
		panic("stride scheduler supports only a single CPU")
	}
	return &stride{
		cpu:     cpus[0],
		queue:   make(job.Jobs, 0),
		quantum: quantum,
	}
}

func (s *stride) Add(job *job.Job) {
	s.queue = append(s.queue, job)
}

// Tick runs the scheduled jobs for the system time, and returns
// the number of jobs finished in this tick. Depending on scheduler requirements,
// the Tick method may assign new jobs to the CPU before returning.
func (s *stride) Tick(systemTime time.Duration) int {
	jobsFinished := 0
	if s.cpu.IsRunning() {
		if s.cpu.Tick() {
			jobsFinished++
		}
	}
	if systemTime%s.quantum == 0 {
		s.reassign()

	}
	return jobsFinished
}

// reassign assigns a job to the cpu
func (s *stride) reassign() {
	if s.cpu.CurrentJob() != nil {
		s.cpu.CurrentJob().Pass += s.cpu.CurrentJob().Stride
		s.Add(s.cpu.CurrentJob())
	}
	nxtJob := s.getNewJob()
	s.cpu.Assign(nxtJob)
}

// getNewJob finds a new job to run on the CPU, removes the job from the queue and returns the job
func (s *stride) getNewJob() *job.Job {
	if len(s.queue) == 0 {
		return nil
	}
	index := MinPass(s.queue)
	removedJob := s.queue[index]
	s.queue = append(s.queue[:index], s.queue[index+1:]...)
	return removedJob
}

// minPass returns the index of the job with the lowest pass value.
func MinPass(theJobs job.Jobs) int {
	lowest := 0
	for i, job := range theJobs {
		if job.Pass < theJobs[lowest].Pass {
			lowest = i
		} else if job.Pass == theJobs[lowest].Pass {
			if job.Stride < theJobs[lowest].Stride {
				lowest = i
			}
		}
	}
	return lowest
}

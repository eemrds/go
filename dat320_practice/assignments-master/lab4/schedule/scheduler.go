package schedule

import "time"

type baseScheduler struct {
	runQueue  chan job
	completed chan result
	jobRunner func(*job)
}

// jobs is a slice of jobs ordered according to some scheduling policies.
type jobs []job

// run starts executing the scheduled jobs from the run queue.
func (s *baseScheduler) run() {
	for i := range s.runQueue {
		latency := time.Now().Sub(i.start)
		s.jobRunner(&i)
		s.completed <- result{i, latency}
	}
	close(s.completed)
}

// results returns the channel of results.
// This is primarily used for testing.
func (s *baseScheduler) results() chan result {
	return s.completed
}

package workerpool

import (
	"net"
	"sync"
)

type Pool struct {
	workers    []*Worker
	jobQueue   chan net.Conn
	maxWorkers int
	wg         sync.WaitGroup
}

// whenever a new pool is created it will create an array of workers
// then each worker in array is put in the job-channel
// the worker.go file handles the Start func. which basically handles the job/task
func NewPool(maxWorkers int) *Pool {

	// init a job-queue/channel which can accomodate 'n' number of connections at a go (n=maxWorkers*2)
	// if more than that comes, we might get errors, so put according to that
	jobQueue := make(chan net.Conn, maxWorkers*2)

	pool := &Pool{
		workers:    make([]*Worker, maxWorkers),
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
	}

	// init workers
	for i := 0; i < maxWorkers; i++ {
		worker := NewWorker(i, jobQueue)
		pool.workers[i] = worker
		pool.wg.Add(1)

		go func(w *Worker) {
			defer pool.wg.Done()
			w.Start()
		}(worker)
	}

	return pool
}

// following func is used to push a connection in the job channel
func (p *Pool) Submit(conn net.Conn) {
	p.jobQueue <- conn
}

func (p *Pool) Stop() {

	// close the job queue
	close(p.jobQueue)

	// stop all workers
	for _, worker := range p.workers {
		worker.Stop()
	}

	// wait for all workers to finish
	p.wg.Wait()
}

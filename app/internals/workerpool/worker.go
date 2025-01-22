package workerpool

import (
	"net"

	"github.com/DivyanshuShekhar55/go-htttpx/app/internals/handler"
)

// the done channel is used to stop/quit the worker pool
// unbuffered channel that signals a message if we need to stop
// see concurrency pattern videos from Kantan Coding
type Worker struct {
	ID       int
	jobQueue chan net.Conn
	done     chan bool
}

func NewWorker(id int, jobQueue chan net.Conn) *Worker {
	return &Worker{
		ID:       id,
		jobQueue: jobQueue,
		done:     make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			select {

			// if a connection is pushed into the queue/channel, take it up
			// pop out one element from the job channel, which is just a request
			case conn := <-w.jobQueue:
				// process the req
				handler.HandleReq(conn)

			case <-w.done:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	// as soon as the stop message passed into queue / channel, the unbuffered channel pushes it to the consumer inside the Start func
	w.done <- true
}

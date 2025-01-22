package workerpool

import "net"

type Worker struct {
	ID int
	jobQueue chan net.Conn
}

func NewWorker(id int, jobQueue chan net.Conn) *Worker{
	return &Worker{
		ID : id,
		jobQueue: jobQueue,
	}
}



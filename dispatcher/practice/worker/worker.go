package worker

import "fmt"

type WorkRequest struct {
	Execute func(config interface{}) error
}

type Worker struct {
	ID         int
	Work       chan WorkRequest
	WorkerPool chan chan WorkRequest
	QuitChan   chan bool
}

type WorkerPoolType chan chan WorkRequest

var WorkQueue = make(chan WorkRequest, 100)

var WorkerPool WorkerPoolType

func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
	return Worker{
		ID:         id,
		Work:       make(chan WorkRequest),
		WorkerPool: workerQueue,
		QuitChan:   make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.Work

			select {
			case work := <-w.Work:
				work.Execute(nil)
			case <-w.QuitChan:
				fmt.Printf("worker %d stopping\n", w.ID)
				return
			}
		}
	}()
}

func ActiveWorkMethod(f func(config interface{}) error) WorkRequest {
	work := WorkRequest{
		Execute: f,
	}
	return work
}

func PushWorkToQueue(w WorkRequest) {
	WorkQueue <- w
}

func SyncWorkerPool(w *WorkerPoolType) {
	WorkerPool = *w
}

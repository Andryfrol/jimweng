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

var (
	WorkQueue  = make(chan WorkRequest, 100)
	WorkerPool WorkerPoolType
	Workers    []*Worker
)

func NewWorker(id int, workerQueue chan chan WorkRequest) *Worker {
	worker := &Worker{
		ID:         id,
		Work:       make(chan WorkRequest),
		WorkerPool: workerQueue,
		QuitChan:   make(chan bool),
	}
	Workers = append(Workers, worker)

	return worker
	// return Worker{
	// 	ID:         id,
	// 	Work:       make(chan WorkRequest),
	// 	WorkerPool: workerQueue,
	// 	QuitChan:   make(chan bool),
	// }
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

// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

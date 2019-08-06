package dispatcher

import (
	"fmt"

	"github.com/jimweng/dispatcher/practice/worker"
)

func StartDispatcher(nworkers int) {
	worker.WorkerPool = make(worker.WorkerPoolType, nworkers)

	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := worker.NewWorker(i+1, worker.WorkerPool)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-worker.WorkQueue:
				fmt.Println("Received work requeust")
				go func() {
					worker := <-worker.WorkerPool

					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}

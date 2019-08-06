package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jimweng/dispatcher/practice/collector"
	"github.com/jimweng/dispatcher/practice/dispatcher"
)

func main() {
	// Start the dispatcher.
	fmt.Println("Starting the dispatcher")
	dispatcher.StartDispatcher(4)

	// Register our collector as an HTTP handler function.
	fmt.Println("Registering the collector")
	http.HandleFunc("/work", collector.Collector)

	// Start the HTTP server!
	go ServerRun()

	// create one chan to print awaiting signal on console
	sigs := make(chan os.Signal, 1)

	// create another chan to receive signal to interrupt original chan
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Println(sig)

		dispatcher.StopWorker()
		time.Sleep(time.Duration(2) * time.Second)

		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}

func ServerRun() {
	fmt.Println("HTTP server listening on", "127.0.0.1:8001")
	if err := http.ListenAndServe("127.0.0.1:8001", nil); err != nil {
		fmt.Println(err.Error())
	}
}

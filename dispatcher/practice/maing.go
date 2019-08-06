package main

import (
	"fmt"
	"net/http"

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
	fmt.Println("HTTP server listening on", "127.0.0.1:8001")
	if err := http.ListenAndServe("127.0.0.1:8001", nil); err != nil {
		fmt.Println(err.Error())
	}
}

package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/goPractice/pkgserver/pkgserver"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

type HelloHandler struct{}

func matric(c chan string, t int) {
	http.Handle("/metrics", promhttp.Handler())

	helloHandler := HelloHandler{}
	http.Handle("/_health", helloHandler)

	log.Fatal(http.ListenAndServe(":8001", nil))
}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rr, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", rr.Message)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("{status:'OK'}")
}

func main() {

	var cc chan string = make(chan string)
	go matric(cc, 5)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Printf("%v\n", sig)

		done <- true
	}()

	<-done
}

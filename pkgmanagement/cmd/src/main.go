package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goPractice/pkgmanagement/config"
	"github.com/goPractice/pkgmanagement/utils"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	action = flag.String("act", "", "please check help doc")
	addr   = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
)

var cfg *config.CrawlerConfig

func main() {
	flag.Parse()

	cfg = config.NewConfig(*action)
	if cfg == nil {
		os.Exit(1)
	}
	var c chan string = make(chan string)

	go matric(c, 5)

	collect()

	// sigs := make(chan os.Signal, 1)
	// done := make(chan bool, 1)
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// go func() {
	// 	sig := <-sigs
	// 	log.Printf("%v\n", sig)

	// 	done <- true
	// }()

	// <-done
}

type HelloHandler struct{}

func matric(c chan string, t int) {
	http.Handle("/metrics", promhttp.Handler())

	helloHandler := HelloHandler{}
	http.Handle("/_health", helloHandler)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/
func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("{status:'OK'}")
}

func collect() {
	var points *[]*utils.PKGContent

	for _, j := range cfg.Inputs {
		fmt.Println(j)
		pts, err := j.Gather()
		if err != nil {
			log.Fatal("%v\n", err)
		}
		points = pts.(*[]*utils.PKGContent)
	}
	// // fmt.Println(points)
	// for i, j := range *points {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	fmt.Printf("%d____the value of j includes Name:%v___Parent:%v___Synopsis:%v___Href:%v\n", i, j.Name, j.Parent, j.Synopsis, j.Href)

	// }

	for i, j := range cfg.Outputs {
		if i == "mysql" {
			if err := j.Write(points); err != nil {
				log.Fatal("%v\n", err)
			}
		}
	}
}

func agent(c chan string, t int) {
	for {
		collect()
		time.Sleep(time.Second * time.Duration(t))
	}
}

func printc(c chan string, t int) {

}

func init() {
}

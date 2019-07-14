package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/goPractice/crawler/crawler/config"
	"github.com/goPractice/crawler/crawler/plugins/inputs/crawler"
	"github.com/goPractice/crawler/crawler/plugins/outputs/mysql"
	"github.com/goPractice/crawler/crawler/utils"
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

	// go agent(c, 15)
	collect()
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

	for i, j := range cfg.Inputs {
		if i == "crawler" {
			var inputCrawler = crawler.QueryUrl{}

			for key, value := range j.(map[string]interface{}) {
				if key == "url" {
					inputCrawler.Url = value.(string)
				}
			}
			cfg.Inputs[i] = inputCrawler
			pts, _ := inputCrawler.Gather()
			points = pts.(*[]*utils.PKGContent)
		}
		// if pts, err := j.(utils.Input).Gather(); err != nil {
		// 	log.Fatal("%v\n", err)
		// } else {
		// 	points = pts.(*[]*utils.PKGContent)
		// }
	}

	for i, j := range cfg.Outputs {
		if i == "mysql" {
			var OutputMysql = mysql.SQLConfig{}

			for key, value := range j.(map[string]interface{}) {
				switch key {
				case "dbname":
					OutputMysql.DBName = value.(string)
				case "dbaddr":
					OutputMysql.DBAddr = value.(string)
				case "password":
					OutputMysql.Password = value.(string)
				case "dbtype":
					OutputMysql.DBType = value.(string)
				case "maxidelconns":
					OutputMysql.MaxIdleConns = int(value.(int64))
				case "maxopenconns":
					OutputMysql.MaxOpenConns = int(value.(int64))
				case "dbport":
					OutputMysql.DBPort = value.(string)
				case "user":
					OutputMysql.User = value.(string)
				case "keepalive":
					OutputMysql.KeepAlive = int(value.(int64))
				}
			}

			if err := OutputMysql.Write(points); err != nil {
				log.Fatal("%v\n", err)
			}
			// if err := j.(utils.Output).Write(points); err != nil {
			// 	log.Fatal("%v\n", err)
			// }
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

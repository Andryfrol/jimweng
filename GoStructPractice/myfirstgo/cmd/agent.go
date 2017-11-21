package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/influxdb/tivan"
	_ "github.com/influxdb/tivan/plugins/all"
)

var fDebug = flag.Bool("debug", false, "show metrics as they're generated to stdout")

var fConfig = flag.String("config", "", "configuration file to load")

func main() {
	flag.Parse()

	var (
		config *tivan.Config
		err    error
	)

	if *fConfig != "" {
		config, err = tivan.LoadConfig(*fConfig)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		config = tivan.DefaultConfig()
	}

	ag := tivan.NewAgent(config)

	if *fDebug {
		ag.Debug = true
	}

	plugins, err := ag.LoadPlugins()
	if err != nil {
		log.Fatal(err)
	}

	shutdown := make(chan struct{})

	signals := make(chan os.Signal)

	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		close(shutdown)
	}()

	log.Print("InfluxDB Agent running")
	log.Printf("Loaded plugins: %s", strings.Join(plugins, " "))
	if ag.Debug {
		log.Printf("Debug: enabled")
		log.Printf("Agent Config: %#v", ag)
	}

	if config.URL != "" {
		log.Printf("Sending metrics to: %s", config.URL)
		log.Printf("Tags enabled: %v", config.ListTags())
	}

	ag.Run(shutdown)
}

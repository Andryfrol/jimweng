package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/jimweng/GoStructPractice/myfirstgo"
	_ "github.com/jimweng/GoStructPractice/myfirstgo/plugins/all"
)

var fDebug = flag.Bool("debug", false, "show metrics as they're generated to stdout")

var fConfig = flag.String("config", "", "configuration file to load")

func main() {
	flag.Parse()

	var (
		config *myfirstgo.Config
		err    error
	)

	if *fConfig != "" {
		config, err = myfirstgo.LoadConfig(*fConfig)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		config = myfirstgo.DefaultConfig()
	}

	fmt.Println(config)
	// ag := myfirstgo.NewAgent(config)

	// if *fDebug {
	// 	ag.Debug = true
	// }

	// plugins, err := ag.LoadPlugins()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	shutdown := make(chan struct{})

	signals := make(chan os.Signal)

	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		close(shutdown)
	}()

	log.Print("InfluxDB Agent running")
	// log.Printf("Loaded plugins: %s", strings.Join(plugins, " "))
	// if ag.Debug {
	// 	log.Printf("Debug: enabled")
	// 	log.Printf("Agent Config: %#v", ag)
	// }

	// if config.URL != "" {
	// 	log.Printf("Sending metrics to: %s", config.URL)
	// 	log.Printf("Tags enabled: %v", config.ListTags())
	// }

	// ag.Run(shutdown)
}

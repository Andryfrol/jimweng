package simpleoutput

// simpleoutput.go

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/outputs"
)

type Simple struct {
	Ok bool
}

func (s *Simple) Description() string {
	return "a demo output"
}

func (s *Simple) SampleConfig() string {
	return "url = localhost"
}

func (s *Simple) Connect() error {
	// Make a connection to the URL here
	return nil
}

func (s *Simple) Close() error {
	// Close connection to the URL here
	return nil
}

func (s *Simple) Write(metrics []telegraf.Metric) error {
	for _, metric := range metrics {
		// write `metric` to the output sink here
	}
	return nil
}

func init() {
	outputs.Add("simpleoutput", func() telegraf.Output { return &Simple{} })
}

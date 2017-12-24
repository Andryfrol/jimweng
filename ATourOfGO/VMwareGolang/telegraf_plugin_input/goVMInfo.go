package main

// simple.go
import (
	"fmt"

	"github.com/goPractice/ATourOfGO/VMwareGolang/telegraf_plugin_input/example"
)

// "github.com/influxdata/telegraf"
// "github.com/influxdata/telegraf/plugins/inputs"

type Simple struct {
	Ok bool
}

// func (s *Simple) Description() string {
// 	return "a demo plugin"
// }

// func (s *Simple) SampleConfig() string {
// 	return "ok = true # indicate if everything is fine"
// }

func (s *Simple) Gather(acc example.Accumulator) error {
	if s.Ok {
		acc.AddFields("state", map[string]interface{}{"value": "pretty good"}, nil)
	} else {
		acc.AddFields("state", map[string]interface{}{"value": "not great"}, nil)
	}

	return nil
}

// func init() {
// 	inputs.Add("simple", func() telegraf.Input { return &Simple{} })
// }
func main() {
	fmt.Println("his")
}

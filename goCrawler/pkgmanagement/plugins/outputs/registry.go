package outputs

import "github.com/mlytics/micro-reporter/utils"

type Creator func() utils.Output

var Outputs = map[string]Creator{}

func Add(name string, creator Creator) {
	Outputs[name] = creator
}

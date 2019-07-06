package inputs

import "github.com/goPractice/pkgmanagement/utils"

type Creator func() utils.Input

var Inputs = map[string]Creator{}

func Add(name string, creator Creator) {
	Inputs[name] = creator
}
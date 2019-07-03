package inputs

import "github.com/goPractice/goCrawler/pkgmanagement/utils"

type Creator func() utils.Input

var Inputs = map[string]Creator{}

func Add(name string, creator Creator) {
	Inputs[name] = creator
}

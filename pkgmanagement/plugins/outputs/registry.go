package outputs

import "github.com/goPractice/goCrawler/pkgmanagement/utils"

type Creator func() utils.Output

var Outputs = map[string]Creator{}

func Add(name string, creator Creator) {
	Outputs[name] = creator
}

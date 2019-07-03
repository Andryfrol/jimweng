package crawler

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/goPractice/goCrawler/pkgmanagement/plugins/inputs"
	"github.com/goPractice/goCrawler/pkgmanagement/utils"
)

type QueryUrl struct {
	Url string
}

type PkgNode struct {
	Name     string
	Synopsis string
	Href     string
}

func (q *QueryUrl) Gather() (interface{}, error) {
	doc, err := goquery.NewDocument(q.Url)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	PkgList, err := parseDoc(doc)
	return PkgList, err
}

func stripTaps(str string) string {
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}

func parseDoc(doc *goquery.Document) (*[]*PkgNode, error) {
	var PkgList []*PkgNode

	/*
		DOM: [table]
		~ ~ ~
		|-tr
		|	|---td
		|	|	|--cls `pkg-name`
		|	|		|-a
		|	|
		|	|---td
		|		|--cls `pkg-synopsis`
		~ ~ ~
	*/
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			pkgnode := PkgNode{}

			rowhtml.Find(".pkg-name").Each(func(indexPkgName int, tableClsPkgName *goquery.Selection) {
				pkgnode.Name = stripTaps(tableClsPkgName.Text())

				tableClsPkgName.Find("a").Each(func(indexa int, tdcell *goquery.Selection) {
					href_text, ok := tdcell.Attr("href")
					if ok {
						pkgnode.Href = stripTaps(href_text)
					}
				})
			})

			rowhtml.Find(".pkg-synopsis").Each(func(indexSynpopsis int, tableClsSynopsis *goquery.Selection) {
				pkgnode.Synopsis = stripTaps(tableClsSynopsis.Text())
			})

			PkgList = append(PkgList, &pkgnode)

		})
	})
	return &PkgList, nil
}

func init() {
	inputs.Add("crawler", func() utils.Input {
		return &QueryUrl{}
	})
}

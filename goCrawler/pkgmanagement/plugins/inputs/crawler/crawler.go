package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type QueryUrl struct {
	Url string
}

type PkgNode struct {
	Name     string
	Synopsis string
	Href     string
}

func main() {
	doc, err := goquery.NewDocument("https://golang.google.cn/pkg/")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

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
				pkgnode.Name = striptaps(tableClsPkgName.Text())

				tableClsPkgName.Find("a").Each(func(indexa int, tdcell *goquery.Selection) {
					href_text, ok := tdcell.Attr("href")
					if ok {
						pkgnode.Href = striptaps(href_text)
					}
				})
			})

			rowhtml.Find(".pkg-synopsis").Each(func(indexSynpopsis int, tableClsSynopsis *goquery.Selection) {
				pkgnode.Synopsis = striptaps(tableClsSynopsis.Text())
			})

			PkgList = append(PkgList, &pkgnode)

		})
	})

	for i, j := range PkgList {

		fmt.Printf("%d\n%v\n%v\n%v\n----\n", i, j.Href, j.Name, j.Synopsis)
	}
}

func striptaps(str string) string {
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}

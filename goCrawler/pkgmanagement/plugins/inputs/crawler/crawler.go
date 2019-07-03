package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type PkgNode struct {
	Name     string
	Synopsis string
	Href     string
	// Child    *PkgNode
}

func main() {
	doc, err := goquery.NewDocument("https://golang.google.cn/pkg/")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	var PkgList []*PkgNode

	// 尋找table，在這個case只有pkg table
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		// table 下每一個row即為一個pkg管理套件
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("td").Each(func(indextd int, tablecell *goquery.Selection) {
				pkgnode := PkgNode{}
				pkgnode.Name = tablecell.Text()
				// row = append(row, tablecell.Text())
				tablecell.Find("a").Each(func(indexa int, tdcell *goquery.Selection) {
					href_text, ok := tdcell.Attr("href")
					if ok {
						pkgnode.Href = href_text
						// row_href = append(row_href, href_text)
					}
				})
				PkgList = append(PkgList, &pkgnode)
			})

		})
	})

	for i, j := range PkgList {
		fmt.Printf("%d___%v___%v__%v\n", i, j.Href, j.Name, j.Synopsis)
	}
}

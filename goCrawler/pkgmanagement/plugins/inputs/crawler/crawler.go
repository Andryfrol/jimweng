package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type PkgNode struct {
	Name     string
	Synopsis string
	Href     string
	Child    *PkgNode
}

func main() {
	// var headings, row, row_href []string
	var row, row_href []string
	var rows, row_hrefs [][]string
	doc, err := goquery.NewDocument("https://golang.google.cn/pkg/")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	// var PkgList []*PkgNode
	// var pkgnode *PkgNode

	// 尋找table，在這個case只有pkg table
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		// table 下每一個row即為一個pkg管理套件
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			// rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {
			// headings = append(headings, tableheading.Text())
			// })
			rowhtml.Find("td").Each(func(indextd int, tablecell *goquery.Selection) {
				// pkgnode.Name = tablecell.Text()
				row = append(row, tablecell.Text())
				tablecell.Find("a").Each(func(indexa int, tdcell *goquery.Selection) {
					href_text, ok := tdcell.Attr("href")
					if ok {
						row_href = append(row_href, href_text)
					}
				})

			})

			rows = append(rows, row)
			row = nil

			row_hrefs = append(row_hrefs, row_href)
			row_href = nil
		})
	})
	// fmt.Println("####### headings = ", len(headings), headings)
	fmt.Println("####### rows = ", len(rows), rows)
	fmt.Println("####### href = ", len(row_hrefs), row_hrefs)
}

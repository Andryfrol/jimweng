package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var headings, row, row_href []string
	var rows, row_hrefs [][]string
	doc, err := goquery.NewDocument("https://golang.google.cn/pkg/")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {
				headings = append(headings, tableheading.Text())
			})
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				row = append(row, tablecell.Text())
			})

			rows = append(rows, row)
			row = nil

			rowhtml.Find("a").Each(func(indexth int, tablehref *goquery.Selection) {
				href_text, ok := tablehref.Attr("href")
				if ok {
					row_href = append(row_href, href_text)
				}
			})
			row_hrefs = append(row_hrefs, row_href)
			row_href = nil
		})
	})
	fmt.Println("####### headings = ", len(headings), headings)
	fmt.Println("####### rows = ", len(rows), rows)
	fmt.Println("####### href = ", len(row_hrefs), row_hrefs)
}

package scraper

import "github.com/gocolly/colly"

func Start() {
	c := colly.NewCollector()

	var blogs []Blog

	c.OnHTML(".blogtitle", func(e *colly.HTMLElement) {
		title := e.ChildText("a")
		date := e.ChildText(".date")
		author := e.ChildText(".author")
		blogs := append(blogs, Blog{Title: title, Date: date, Author: author})
	})
}

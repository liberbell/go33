package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("#blogindex", func(e *colly.HTMLElement) {
		date := e.ChildText("p:nth-child(1) > span.date")
		fmt.Println("Date: ", date)

		title := e.ChildText("p:nth-child(1) > a")
		fmt.Println("Title: ", title)

		author := e.ChildText("p:nth-child(1) > span.author")
		fmt.Println("Author: ", author)
	})

	err := c.Visit("https://go.dev/blog")
	if err != nil {
		log.Fatal(err)
	}
}

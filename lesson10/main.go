package main

import "scraper/scraper"

func main() {
	scraper.InitDB()
	scraper.Start()
}

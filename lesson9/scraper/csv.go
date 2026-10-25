package scraper

import (
	"encoding/csv"
	"os"
)

func SaveToCsv(blog Blog) {
	file, err := os.OpenFile("blogs.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
}

package main

import (
	"github.com/TheusLab/ASN-Project/backend/utils"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		utils.Log.Info().Str("link", link).Msg("Link found")
		e.Request.Visit(link)
	})

	c.Visit("https://example.com")
}

package main

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly/v2"
)

func main() {

	urlFilters := regexp.MustCompile(`dontstarve.fandom.com/wiki/[^(Special:Log)|
	^(Template)|^(Help)|^(User)]`)

	c := colly.NewCollector(
		colly.AllowedDomains("dontstarve.fandom.com", "www.dontstarve.fandom.com"),
		colly.URLFilters(urlFilters),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://dontstarve.fandom.com/wiki/Magma")

}

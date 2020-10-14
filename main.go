package main

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/gocolly/colly/v2"

	"github.com/KalilCazes/dontstarve-crawler/character"
)

//NewCollector creates a collector with predefined configuration
func NewCollector() *colly.Collector {

	urlFilters := regexp.MustCompile(`dontstarve.fandom.com/wiki/[^(Special:Log)|
	^(Template)|^(Help)|^(User)]`)

	c := colly.NewCollector(
		colly.AllowedDomains("dontstarve.fandom.com", "www.dontstarve.fandom.com"),
		colly.URLFilters(urlFilters),
	)
	err := c.Limit(&colly.LimitRule{
		DomainGlob: "dontstarve.fandom.com*",
		Delay:      1 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	return c
}
func main() {
	collector := NewCollector()
	chars := character.GetCharacters(collector)
	for _, name := range chars {
		fmt.Printf("%s: %s\n", name, character.GetInfo(collector, name).Nickname)
	}
}

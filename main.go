package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type character struct {
	name         string
	nickname     string
	motto        string
	bio          string
	perk         string
	profileImage string
	health       int
	hunger       int
	sanity       int
	favoriteFood map[string]string
}

func main() {

	var characters []character

	urlFilters := regexp.MustCompile(`dontstarve.fandom.com/wiki/[^(Special:Log)|
	^(Template)|^(Help)|^(User)]`)

	c := colly.NewCollector(
		colly.AllowedDomains("dontstarve.fandom.com", "www.dontstarve.fandom.com"),
		colly.URLFilters(urlFilters),
	)
	

	c.OnHTML("center", func(e *colly.HTMLElement) {

		e.ForEach("tbody", func(_ int, tbody *colly.HTMLElement) {
			rowIdentifier := 0

			e.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
				characterIndex := 0
				tr.ForEach("td", func(_ int, td *colly.HTMLElement) {
					content := strings.TrimSpace(td.Text)

					switch rowIdentifier % 3 {
					case 0:
						var c character
						c.profileImage = content
						characters = append(characters, c)
						// fmt.Printf("[0] characters: %#v\n", characters)
					case 1:
						// fmt.Printf("[1] characters: %#v\n", characters)
						characters[characterIndex].name = content
					case 2:
						// fmt.Printf("[2] characters: %#v\n", characters)
						characters[characterIndex].perk = content
					}
					characterIndex++

				})
				rowIdentifier++
				characterIndex = 0
				// write to file for future parsing
				filePath := fmt.Sprintf("html/%s", filepath.Base(e.Request.URL.Path))
				err := ioutil.WriteFile(filePath, []byte(e.Text), 0644)
				if err != nil {
					fmt.Printf("%#v", err)
				}
			})
		})
		fmt.Printf("characters: %#v\n", characters)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://dontstarve.fandom.com/wiki/Characters")
	fmt.Println("O")
}

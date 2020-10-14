package character

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

//Character struct containing information about character
type Character struct {
	Name         string
	Nickname     string
	Motto        string
	Bio          string
	Perk         []string
	ProfileImage string
	Health       int
	Hunger       int
	Sanity       int
	FavoriteFood map[string]string
}

//GetCharacters returns a slice of characters name
func GetCharacters(c *colly.Collector) []string {

	var charactersName []string

	c.OnHTML("center", func(e *colly.HTMLElement) {

		charSelector := "table > tbody > tr > td > b > a"

		e.ForEach(charSelector, func(index int, item *colly.HTMLElement) {

			charactersName = append(charactersName, item.Text)

		})

	})

	c.Visit("https://dontstarve.fandom.com/wiki/Characters")
	return charactersName
}

//GetInfo get information about specific character
func GetInfo(c *colly.Collector, characterName string) Character {

	character := Character{}
	c.OnHTML("div.pi-section-content", func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		nickname := goquerySelection.Find("[data-source=\"nick dst\"] .pi-data-value").Text()
		character.Nickname = nickname
	})

	c.Visit(fmt.Sprintf("https://dontstarve.fandom.com/wiki/%s", characterName))
	return character
}

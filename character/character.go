package character

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

//Character struct containing information about character
type Character struct {
	Name         string
	Nickname     string
	Motto        string
	Bio          string
	Perks        []string
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

	err := c.Visit("https://dontstarve.fandom.com/wiki/Characters")
	checkError(err)

	return charactersName
}

//GetInfo get information about specific character
func GetInfo(c *colly.Collector, characterName string) Character {

	character := Character{}
	c.OnHTML("div.pi-section-content:nth-child(2)", func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		nickname := goquerySelection.Find("[data-source=\"nick dst\"] .pi-data-value").Text()
		character.Nickname = nickname
	})

	c.OnHTML("div.pi-section-content:nth-child(2)", func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		motto := goquerySelection.Find("[data-source=\"motto dst\"] .pi-data-value").Text()
		character.Motto = motto
	})

	c.OnHTML("div.pi-section-content:nth-child(2)", func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		bio := goquerySelection.Find("[data-source=\"bio\"] .pi-data-value").Text()
		character.Bio = bio
	})

	c.OnHTML("div.pi-section-content:nth-child(2)", func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		s, err := goquerySelection.Find("[data-source=\"perk dst\"] .pi-data-value").Html()
		checkError(err)

		perks := strings.Split(s, "<br/>")
		character.Perks = perks
	})

	c.OnHTML("div.pi-section-content:nth-child(2)", func(e *colly.HTMLElement) {
		var err error
		pi := e.ChildAttr("section:nth-child(1) > figure:nth-child(1) > a:nth-child(1) > img", "src")
		pi, err = trimImageURL(pi)
		checkError(err)

		character.ProfileImage = pi
	})

	c.OnHTML("div.pi-section-content:nth-child(2)", func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		ht := goquerySelection.Find("div.pi-section-content:nth-child(2) > section:nth-child(2) > table:nth-child(1) > tbody:nth-child(3) > tr:nth-child(1) > td:nth-child(1)").Text()
		health, err := strconv.Atoi(ht)
		checkError(err)

		character.Health = health
	})

	c.OnHTML("div.pi-section-content:nth-child(2)", func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		hg := goquerySelection.Find("div.pi-section-content:nth-child(2) > section:nth-child(2) > table:nth-child(1) > tbody:nth-child(3) > tr:nth-child(1) > td:nth-child(2)").Text()
		hunger, err := strconv.Atoi(hg)
		checkError(err)

		character.Hunger = hunger
	})

	c.OnHTML("div.pi-section-content:nth-child(2)", func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		s := goquerySelection.Find("div.pi-section-content:nth-child(2) > section:nth-child(2) > table:nth-child(1) > tbody:nth-child(3) > tr:nth-child(1) > td:nth-child(3)").Text()
		sanity, err := strconv.Atoi(s)
		checkError(err)

		character.Sanity = sanity
	})

	err := c.Visit(fmt.Sprintf("https://dontstarve.fandom.com/wiki/%s", characterName))
	checkError(err)

	return character
}

func trimImageURL(url string) (string, error) {
	s := strings.SplitAfter(url, "png")
	if len(s) < 1 {
		return "", fmt.Errorf("Invalid Image URL: %s", url)
	}
	return s[0], nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

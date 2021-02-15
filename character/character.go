package character

import (
	"fmt"
	"log"
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
	Health       string
	Hunger       string
	Sanity       string
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

//GetCharacterInfo get information about specific character
func GetCharacterInfo(c *colly.Collector, characterName string) Character {
	var filter string

	c.OnHTML("div.pi-section-contents [data-ref]", func(e *colly.HTMLElement) {

		var tabSize string
		tabSize = e.Attr("data-ref")

		if tabSize == "1" {
			filter = ":nth-child(2)"
		} else {
			filter = ":nth-child(1)"
		}

	})
	character := Character{}
	c.OnHTML("div.pi-section-content"+filter, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		nickname := goquerySelection.Find("section.pi-item:nth-child(1) > div:nth-child(2) > div:nth-child(2)").Text()
		character.Nickname = nickname
	})

	c.OnHTML("div.pi-section-content"+filter, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		motto := goquerySelection.Find("[data-source=\"motto dst\"] .pi-data-value").Text()

		character.Motto = motto
	})

	c.OnHTML("div.pi-section-content"+filter, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		bio := goquerySelection.Find("[data-source=\"bio\"] .pi-data-value").Text()
		character.Bio = bio
	})

	c.OnHTML("div.pi-section-content"+filter, func(e *colly.HTMLElement) {

		goquerySelection := e.DOM
		s, err := goquerySelection.Find("[data-source=\"perk dst\"] .pi-data-value").Html()
		checkError(err)

		perks := trimPerk(s)

		character.Perks = perks
	})

	c.OnHTML("div.pi-section-content"+filter, func(e *colly.HTMLElement) {
		var err error
		pi := e.ChildAttr("section:nth-child(1) > figure:nth-child(1) > a:nth-child(1) > img", "src")
		pi, err = trimImageURL(pi)
		checkError(err)

		character.ProfileImage = pi
	})

	c.OnHTML("div.pi-section-content"+filter, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		health := goquerySelection.Find("section:nth-child(2) > table:nth-child(1) > tbody:nth-child(3) > tr:nth-child(1) > td:nth-child(1)").Text()

		character.Health = health
	})

	c.OnHTML("div.pi-section-content"+filter, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		hunger := goquerySelection.Find("section:nth-child(2) > table:nth-child(1) > tbody:nth-child(3) > tr:nth-child(1) > td:nth-child(2)").Text()

		character.Hunger = hunger
	})

	c.OnHTML("div.pi-section-content"+filter, func(e *colly.HTMLElement) {
		goquerySelection := e.DOM
		sanity := goquerySelection.Find("section:nth-child(2) > table:nth-child(1) > tbody:nth-child(3) > tr:nth-child(1) > td:nth-child(3)").Text()

		character.Sanity = sanity
	})

	c.OnHTML("div.pi-section-content"+filter, func(e *colly.HTMLElement) {
		var l, foodName []string
		var link string
		food := "section [data-source=\"favorite food\"] > div"
		e.ForEach(food, func(index int, item *colly.HTMLElement) {

			foodName = item.ChildAttrs("a", "title")
			l = item.ChildAttrs("a", "href")

		})

		character.FavoriteFood = make(map[string]string)
		for i := 0; i < len(foodName); i++ {

			link = "https://dontstarve.fandom.com" + l[i]
			character.FavoriteFood[foodName[i]] = link
		}
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

func trimPerk(raw string) []string {
	var aux string
	var copy bool = true
	tmp := strings.Replace(raw, "<br/>", "\n", -1)

	for _, c := range tmp {

		switch c {
		case '<':
			copy = false

		case '>':
			copy = true
			continue
		}

		if copy == true {
			aux += string(c)
		}

	}

	p := strings.Split(aux, "\n")

	return p
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

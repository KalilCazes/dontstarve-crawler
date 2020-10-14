package character

import (
	"regexp"
	"testing"
	"time"

	"github.com/gocolly/colly/v2"
)

func createCollector() *colly.Collector {
	urlFilters := regexp.MustCompile(`dontstarve.fandom.com/wiki/[^(Special:Log)|
	^(Template)|^(Help)|^(User)]`)

	c := colly.NewCollector(
		colly.AllowedDomains("dontstarve.fandom.com", "www.dontstarve.fandom.com"),
		colly.URLFilters(urlFilters),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob: "dontstarve.fandom.com*",
		Delay:      1 * time.Second,
	})

	return c
}
func TestGetInfo(t *testing.T) {
	c := createCollector()
	tests := []Character{
		{
			Name:         "Maxwell",
			Nickname:     "The Puppet Master",
			Motto:        "Freedom!",
			Perk:         []string{"Is dapper but frail.", "Can split his mind into pieces.", "On a first-name basis with the night."},
			Health:       75,
			Hunger:       150,
			Sanity:       200,
			FavoriteFood: map[string]string{"Wobster_Dinner": "https://dontstarve.fandom.com/wiki/Wobster_Dinner"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			expected := tt.Nickname
			got := GetInfo(c, tt.Name).Nickname

			if got != expected {
				t.Fatalf("%s: expected: %v, got: %v", tt.Nickname, expected, got)
			}
		})
	}
}

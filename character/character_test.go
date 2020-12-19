package character

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gocolly/colly/v2"
)

func createCollector() *colly.Collector {

	fake := httptest.NewServer(http.FileServer(http.Dir(".")))

	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)
	c.Visit(fake.URL)
	return c
}

func TestGetInfo(t *testing.T) {
	c := createCollector()

	tests := []Character{
		{
			Name:         "Maxwell",
			Nickname:     "The Puppet Master",
			Motto:        "\"Freedom suits me.\"",
			Bio:          "Formerly the Shadow King, lately Maxwell finds himself reacquainted with life among the commonfolk.",
			Perks:        []string{"Is dapper, but frail", "Can split his mind into pieces", "Was once the king of the world"},
			ProfileImage: "https://static.wikia.nocookie.net/dont-starve-game/images/9/95/Maxwell_DST.png",
			Health:       75,
			Hunger:       150,
			Sanity:       200,
			FavoriteFood: map[string]string{"Wobster_Dinner": "https://dontstarve.fandom.com/wiki/Wobster_Dinner"},
		},
	}

	for _, tt := range tests {

		t.Run("get nickname information", func(t *testing.T) {
			expected := tt.Nickname
			got := GetInfo(c, tt.Name).Nickname

			if got != expected {
				t.Errorf("%s: expected: %v, got: %v", "[nickname]", expected, got)
			}
		})

		t.Run("get motto information", func(t *testing.T) {
			expected := tt.Motto
			got := GetInfo(c, tt.Name).Motto

			if got != expected {
				t.Errorf("%s: expected: %v, got: %v", "[motto]", expected, got)
			}
		})

		t.Run("get bio information", func(t *testing.T) {
			expected := tt.Bio
			got := GetInfo(c, tt.Name).Bio

			if got != expected {
				t.Errorf("%s: expected: %v, got: %v", "[bio]", expected, got)
			}
		})

		t.Run("get perk information", func(t *testing.T) {
			expected := tt.Perks
			got := GetInfo(c, tt.Name).Perks

			if !reflect.DeepEqual(got, expected) {
				t.Errorf("%s: expected: %v, got: %v", "[perks]", expected, got)
			}
		})

		t.Run("get profile image information", func(t *testing.T) {
			expected := tt.ProfileImage
			got := GetInfo(c, tt.Name).ProfileImage

			if got != expected {
				t.Errorf("%s: expected: %v, got: %v", "[profile image]", expected, got)
			}
		})

		t.Run("get health information", func(t *testing.T) {
			expected := tt.Health
			got := GetInfo(c, tt.Name).Health

			if got != expected {
				t.Errorf("%s: expected: %v, got: %v", "[health]", expected, got)
			}
		})

		t.Run("get hunger information", func(t *testing.T) {
			expected := tt.Hunger
			got := GetInfo(c, tt.Name).Hunger

			if got != expected {
				t.Errorf("%s: expected: %v, got: %v", "[hunger]", expected, got)
			}
		})

		t.Run("get sanity information", func(t *testing.T) {
			expected := tt.Sanity
			got := GetInfo(c, tt.Name).Sanity

			if got != expected {
				t.Errorf("%s: expected: %v, got: %v", "[sanity]", expected, got)
			}
		})
	}
}

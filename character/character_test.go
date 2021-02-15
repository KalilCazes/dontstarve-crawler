package character

import (
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gocolly/colly/v2"
)

func createCollector() *colly.Collector {

	fake := httptest.NewServer(http.FileServer(http.Dir("test-data")))

	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)
	err := c.Visit(fake.URL)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func TestGetCharacterInfo(t *testing.T) {
	c := createCollector()

	tests := []Character{
		{
			Name:         "Maxwell",
			Nickname:     "The Puppet Master",
			Motto:        "\"Freedom suits me.\"",
			Bio:          "Formerly the Shadow King, lately Maxwell finds himself reacquainted with life among the commonfolk.",
			Perks:        []string{"Is dapper, but frail", "Can split his mind into pieces", "Was once the king of the world"},
			ProfileImage: "https://static.wikia.nocookie.net/dont-starve-game/images/9/95/Maxwell_DST.png",
			Health:       "75",
			Hunger:       "150",
			Sanity:       "200",
			FavoriteFood: map[string]string{"Wobster Dinner (DST)": "https://dontstarve.fandom.com/wiki/Wobster_Dinner_(DST)"},
		},

		{
			Name:         "Wortox",
			Nickname:     "The Soul Starved",
			Motto:        "\"Pardon me if I don't shake your hand.\"",
			Bio:          "Wortox is a mischievous imp whose levity masks a deep well of guilt. He has taken many souls in his time.",
			Perks:        []string{"Is an imp", "Can hop through time and space", "Less nourished by physical food"},
			ProfileImage: "https://static.wikia.nocookie.net/dont-starve-game/images/6/6f/Wortox_DST.png",
			Health:       "200",
			Hunger:       "175",
			Sanity:       "150",
			FavoriteFood: map[string]string{"Pomegranate": "https://dontstarve.fandom.com/wiki/Pomegranate",
				"Sliced Pomegranate": "https://dontstarve.fandom.com/wiki/Sliced_Pomegranate"},
		},

		{
			Name:         "Wilson",
			Nickname:     "The Gentleman Scientist",
			Motto:        "\"I will conquer this world with the power of my MIND!\"",
			Bio:          "At heart Wilson is truly a man of science. It's not in his nature to resist the call of the unknown.",
			Perks:        []string{"Grows a magnificent beard"},
			ProfileImage: "https://static.wikia.nocookie.net/dont-starve-game/images/3/33/Wilson_DST.png",
			Health:       "150",
			Hunger:       "150",
			Sanity:       "200",
			FavoriteFood: map[string]string{"Bacon and Eggs": "https://dontstarve.fandom.com/wiki/Bacon_and_Eggs"},
		},
	}

	for _, tt := range tests {

		t.Run("get nickname information", func(t *testing.T) {
			expected := tt.Nickname
			got := GetCharacterInfo(c, tt.Name).Nickname

			if got != expected {
				t.Errorf("%s: expected: %v, got: %v", "[nickname]", expected, got)
			}
		})

		t.Run("get motto information", func(t *testing.T) {
			expected := tt.Motto
			got := GetCharacterInfo(c, tt.Name).Motto

			if got != expected {
				t.Errorf("%s: expected: %v, got: %v", "[motto]", expected, got)
			}
		})

		t.Run("get bio information", func(t *testing.T) {
			expected := tt.Bio
			got := GetCharacterInfo(c, tt.Name).Bio

			if got != expected {
				t.Errorf("%s: expected: %v, got: %v", "[bio]", expected, got)
			}
		})

		t.Run("get perk information", func(t *testing.T) {
			expected := tt.Perks
			got := GetCharacterInfo(c, tt.Name).Perks

			if !reflect.DeepEqual(got, expected) {
				t.Errorf("%s: expected: %v, got: %v", "[perks]", expected, got)
			}
		})

		t.Run("get profile image information", func(t *testing.T) {
			expected := tt.ProfileImage
			got := GetCharacterInfo(c, tt.Name).ProfileImage

			if got != expected {
				t.Errorf("%s: expected: %v, got: %v", "[profile image]", expected, got)
			}
		})

		t.Run("get health information", func(t *testing.T) {
			expected := tt.Health
			got := GetCharacterInfo(c, tt.Name).Health

			if got != expected {
				t.Errorf("%s: expected: %v, got: %v", "[health]", expected, got)
			}
		})

		t.Run("get hunger information", func(t *testing.T) {
			expected := tt.Hunger
			got := GetCharacterInfo(c, tt.Name).Hunger

			if got != expected {
				t.Errorf("%s: expected: %v, got: %v", "[hunger]", expected, got)
			}
		})

		t.Run("get sanity information", func(t *testing.T) {
			expected := tt.Sanity
			got := GetCharacterInfo(c, tt.Name).Sanity

			if got != expected {
				t.Errorf("%s: expected: %v, got: %v", "[sanity]", expected, got)
			}
		})

		t.Run("get favorite food information", func(t *testing.T) {
			expected := tt.FavoriteFood
			got := GetCharacterInfo(c, tt.Name).FavoriteFood

			if !reflect.DeepEqual(got, expected) {
				t.Errorf("%s: expected: %v, got: %v", "[favorite food]", expected, got)
			}
		})
	}
}

package scrapers

import (
	"fmt"
	"os"
	"time"

	"github.com/HugoJBello/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TwitterScraper interface {
	Load()
	ScrapDate() []interface{}
}

type TwitterScraperSuicide struct {
	Client *twitter.Client
}

func (scraper *TwitterScraperSuicide) Load() {
	consumerKey := os.Getenv("API_KEY")
	consumerSecret := os.Getenv("API_TOKEN")
	accessToken := os.Getenv("ACCESS_TOKEN_KEY")
	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	scraper.Client = twitter.NewClient(httpClient)
}

func (scraper TwitterScraperSuicide) ScrapDate(date time.Time) []twitter.Tweet {
	formatedDate := date.Format("2006-01-02")
	dateBefore := date.AddDate(0, 0, -1)
	formatedDateBefore := dateBefore.Format("2006-01-02")
	fmt.Printf("searching twitts from %s to %s", formatedDateBefore, formatedDate)
	search, _, err := scraper.Client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "#suicidio", Geocode: "42.4,-3.7,650km", Until: formatedDate, Since: formatedDateBefore})

	for _, twitt := range search.Statuses {
		fmt.Println(twitt)
	}
	// fmt.Println(search.Statuses)
	// fmt.Println(resp)
	fmt.Println(err)
	return search.Statuses
}

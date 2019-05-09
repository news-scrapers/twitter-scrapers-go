package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"ithub.com/dghubble/go-twitter/twitter"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	costumerKey := ""
	consumerSecret := ""
	accessToken := ""
	accessSecret := ""

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
	tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})

	// Send a Tweet
	tweet, resp, err := client.Statuses.Update("just setting up my twttr", nil)

	// Status Show
	tweet, resp, err := client.Statuses.Show(585613041028431872, nil)

	// Search Tweets
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "gopher",
	})

	// User Show
	user, resp, err := client.Users.Show(&twitter.UserShowParams{
		ScreenName: "dghubble",
	})

	// Followers
	followers, resp, err := client.Followers.List(&twitter.FollowerListParams{})

}

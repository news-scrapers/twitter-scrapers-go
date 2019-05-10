package main

import (
	"fmt"
	"os"

	"github.com/HugoJBello/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	consumerKey := os.Getenv("API_KEY")
	consumerSecret := os.Getenv("API_TOKEN")
	accessToken := os.Getenv("ACCESS_TOKEN_KEY")
	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET_SECRET_SECRET_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// //#suicidio geocode:42.4,-3.7,1000km since:2019-05-01 until:2019-05-02 count=1000
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "#suicidio", Until: "2019-05-02", Since: "2019-05-01"})

	fmt.Println(search.Statuses)
	// fmt.Println(resp)
	fmt.Println(err)

}

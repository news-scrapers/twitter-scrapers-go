package scrapers

import (
	"fmt"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestScraper(t *testing.T) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	scraper := TwitterScraperSuicide{}
	scraper.Load()

	date := time.Date(2019, 05, 04, 0, 0, 0, 0, time.UTC)
	// date := time.Now()

	twitts := scraper.ScrapDate(date)
	fmt.Println(twitts)

}

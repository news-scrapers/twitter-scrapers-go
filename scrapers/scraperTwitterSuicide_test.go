package scrapers

import (
	"fmt"
	"testing"
	"time"
)

func TestScraper(t *testing.T) {
	scraper := TwitterScraperSuicide{}
	scraper.Load()

	twitts := scraper.ScrapDate(time.Now())
	fmt.Println(twitts)

}

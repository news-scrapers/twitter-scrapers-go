package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//colly.AllowedDomains("twitter.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("div", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://twitter.com/search?q=%23suicidio%20geocode%3A42.4%2C-3.7%2C650km%20since%3A2019-05-01%20until%3A2019-05-05&src=typd")
}

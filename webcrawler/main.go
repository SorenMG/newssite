// Webcrawler

// Stuff to think about:
// Dynamic websites - browser automation (selenium)
// Queue
// Concurrency
// One collector per site scraped
// Persistent background storage
// Cross referencing filter with background storage (Maybe store the most frequent mentioned links in memory)
// Scraping article metadata
// Keep scraping in scope of domain

package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domainsr hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.bt.dk"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping
	c.Visit("https://www.instagram.com/marinaasmussen/")
}

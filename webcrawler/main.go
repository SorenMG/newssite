// Webcrawler

// Stuff to think about:
// Dynamic websites - browser automation (selenium)
// Queue - Min priority queue then count number of times link occur
// Concurrency
// One collector per site scraped
// Persistent background storage
// Cross referencing filter with background storage (Maybe store the most frequent mentioned links in memory)
// Scraping article metadata
// Keep scraping in scope of domain
// Cookies
// Fix relative paths to absolute paths

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Missing URL argument")
		os.Exit(1)
	}

	PageURL := os.Args[1]

	// Instantiate default collector
	c := colly.NewCollector(
	//colly.AllowedDomains("www.bt.dk"),
	)
	c.DisableCookies()
	c.Limit(&colly.LimitRule{Parallelism: 4})

	// create a request queue with 2 consumer threads
	q, _ := queue.New(
		2, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	// Check article metadata
	c.OnHTML("head", func(e *colly.HTMLElement) {
		//fmt.Println("Head found")
	})

	// Queue new links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// Convert relative links to absolute
		if !govalidator.IsURL(link) {
			link = PageURL + link
		}

		hasVisited, err := c.HasVisited(link)

		if err != nil {
			return
		}

		if hasVisited {
			return
		}

		// Visit link found on page on a new thread
		q.AddURL(link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	// Start scraping
	q.AddURL(PageURL)

	q.Run(c)
}

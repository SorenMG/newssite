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
	"log"
	url "net/url"
	"os"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
)

func main() {
	// Check for argument
	if len(os.Args) != 2 {
		log.Println("Missing URL argument")
		os.Exit(1)
	}

	// Parse and validate argument
	pageURL := os.Args[1]
	queryUrl := validateURL(pageURL)

<<<<<<< HEAD
	// Instantiate default collector
	c := colly.NewCollector(
	//colly.AllowedDomains("www.bt.dk"),
	)
	c.DisableCookies()
	c.Limit(&colly.LimitRule{Parallelism: 2})
=======
	// Instantiate collector
	c := initScraper(queryUrl)
>>>>>>> dev

	// create a request queue with 10 consumer threads
	q, _ := queue.New(
		20, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	// Queue new links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))

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
		//log.Println("Visiting:", r.URL.String())
	})

	c.OnHTML("meta[property]", func(e *colly.HTMLElement) {
		if e.Attr("property") != "og:title" {
			return
		}
		log.Println(e.Attr("content"))
	})

	// Start scraping
	q.AddURL(pageURL)

	q.Run(c)
}

func initScraper(url *url.URL) (scraper *colly.Collector) {
	scraper = colly.NewCollector(
		colly.AllowedDomains("cnn.com"),
	)
	scraper.DisableCookies()
	return
}

func validateURL(pageUrl string) (queryUrl *url.URL) {
	queryUrl, err := url.Parse(pageUrl)

	if err != nil {
		log.Println("Could not parse URL")
		os.Exit(1)
	}

	// Validate URL
	if queryUrl.Scheme == "" {
		log.Println("Missing scheme")
		os.Exit(1)
	}

	if queryUrl.Host == "" {
		log.Println("Missing host name")
		os.Exit(1)
	}

	//if !strings.Contains(queryUrl.Host, "www.") {
	//log.Println("Missing www in URL")
	//os.Exit(1)
	//}

	return
}

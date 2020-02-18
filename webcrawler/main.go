/*
Versatile webcrawler for scraping news sites and gather articles.

Made by: Soren Gade
*/

package main

import (
	"fmt"
	"log"
	url "net/url"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/redisstorage"
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

	// Handle html
	c.OnHTML("a[href]", func(e *colly.HTMLElement) { handleLinks(c, e) })
	c.OnHTML("meta[property]", func(e *colly.HTMLElement) { handleMeta(e) })

	// Debug print statement
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting:", r.URL.String())
	})

	// Set start page
	c.Visit(pageURL)
	c.Wait()
}

func handleMeta(e *colly.HTMLElement) {
	if e.Attr("property") != "og:title" {
		return
	}
	//log.Println(e.Attr("content"))
}

func handleLinks(c *colly.Collector, e *colly.HTMLElement) {
	link := e.Request.AbsoluteURL(e.Attr("href"))

	hasVisited, err := c.HasVisited(link)

	if err != nil {
		return
	}

	if hasVisited {
		return
	}

	e.Request.Visit(link)
}

func initScraper(url *url.URL) (scraper *colly.Collector) {
	domainList := strings.Split(url.Host, ".")

	// Take the second and last and make regex
	regular := fmt.Sprintf("(http://|https://)[a-zA-Z]+\\.%s\\.%s", domainList[1], domainList[2])

	// create the redis storage
	storage := &redisstorage.Storage{
		Address:  "127.0.0.1:6379",
		Password: "",
		DB:       0,
		Prefix:   "news_url_filter",
	}

	scraper = colly.NewCollector(
		colly.Async(true),
		colly.URLFilters(
			regexp.MustCompile(regular),
		),
	)
	scraper.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 1})
	scraper.DisableCookies()

	// add storage to the collector
	err := scraper.SetStorage(storage)
	if err != nil {
		panic(err)
	}

	// delete previous data from storage
	if err := storage.Clear(); err != nil {
		log.Fatal(err)
	}

	// close redis client
	//defer storage.Client.Close()

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

	if !strings.Contains(queryUrl.Host, "www.") {
		log.Println("Missing www in URL")
		os.Exit(1)
	}

	return
}

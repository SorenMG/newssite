package main

import (
	"log"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.DisableCookies()

	c.OnHTML("meta", func(e *colly.HTMLElement) {
		if e.Attr("name") != "" {
			log.Println("Meta", e.Attr("name")+":", e.Attr("content"))
		}
		if e.Attr("property") != "" {
			log.Println("Meta", e.Attr("property")+":", e.Attr("content"))
		}
		//log.Println(e)
	})

	log.Println(" -------------------------------------------------------- BT")

	c.Visit("https://www.bt.dk/krimi/forsvundet-kvinde-er-fundet-doed")

	log.Println(" -------------------------------------------------------- CNN")

	c.Visit("https://edition.cnn.com/asia/live-news/coronavirus-outbreak-02-16-20-intl-hnk/index.html")

	log.Println(" -------------------------------------------------------- DR")

	c.Visit("https://www.dr.dk/nyheder/indland/trae-vaeltede-ned-over-mor-og-barn-pige-pa-fire-ar-i-kritisk-tilstand")

	log.Println(" -------------------------------------------------------- TV2")

	c.Visit("https://nyheder.tv2.dk/2020-02-16-trae-vaeltet-i-storm-fireaarig-pige-i-kritisk-tilstand")
}

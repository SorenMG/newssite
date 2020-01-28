package main

import (
    "fmt"
    scraper "webcrawler/scraper"
)

func main() {
    const data = `<!DOCTYPE html>
<html>
<head>
    <title></title>
</head>
<body>
    body content
    <a href="test"></a>
</body>
</html>`
    doc, err := scraper.ParseHTML(data)
    if err != nil {
        return
    }
    links := scraper.GetLinks(doc)
    fmt.Println(links)
}

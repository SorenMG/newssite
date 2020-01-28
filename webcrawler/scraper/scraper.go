package scraper

import (
    "golang.org/x/net/html"
    "strings"
)

func GetLinks(doc *html.Node) []string {
    var crawler func(*html.Node)
    var linkList = []string{}
    crawler = func(node *html.Node) {
        if node.Type == html.ElementNode && node.Data == "a" {
            for i := 0; i < len(node.Attr); i++ {
                if node.Attr[i].Key == "href" {
                    linkList = append(linkList, node.Attr[i].Val)
                }
            }
        }
        for child := node.FirstChild; child != nil; child = child.NextSibling {
            crawler(child)
        }
    }
    crawler(doc)
    return linkList
}

func ParseHTML(data string) (*html.Node, error) {
    return html.Parse(strings.NewReader(data))
}

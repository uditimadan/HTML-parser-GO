package main

import (
	"strings"
	"golang.org/x/net/html"
)

//main function
func main() {

}

//extraction of the words and hrefs from the given HTML body
func extract(body []byte) (words []string, hrefs []string) {
	doc, err := html.Parse(strings.NewReader(string(body))) //parsing the HTML
	if err != nil {
		return
	}
	//recursive function to traverse HTML nodes
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" { //checking for anchor tags
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					hrefs = append(hrefs, attr.Val) //collecting the href attributes
				}
			}
		}
		if n.Type == html.TextNode { 
			words = append(words, strings.Fields(n.Data)...) //text html nodes
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling { //going over the child nodes
			f(c)
		}
	}
	f(doc)
	return
}

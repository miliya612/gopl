package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	counter := make(map[string]int)
	count(counter, doc)
	for k, v := range counter {
		fmt.Printf("%10s: %d\n", k, v)
	}
}

func count(counter map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counter[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(counter, c)
	}
}

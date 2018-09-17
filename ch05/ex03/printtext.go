package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	printText(nil, doc)

}

func printText(stack []string, n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		stack = append(stack, n.Data)
	case html.TextNode:
		if len(stack) == 0 {
			panic("invalid html")
		}
		if stack[len(stack)-1] != "script" && stack[len(stack)-1] != "style" {
			trimmed := strings.TrimSpace(n.Data)
			if len(trimmed) > 0 {
				fmt.Printf("<%s>%s</%[1]s>\n", stack[len(stack)-1], trimmed)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printText(stack, c)
	}
}

func format(n *html.Node) string {
	return n.Data
}

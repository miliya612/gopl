package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)
	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	//n.Data = strings.TrimSpace(n.Data)
	if true {
		switch n.Type {
		case html.CommentNode:
		case html.ElementNode:
			printStartElement(n)
		case html.DoctypeNode:
			printDoctypeElement(n)
		case html.DocumentNode:
			// Do nothing
		case html.TextNode:
			trimmed := strings.TrimSpace(n.Data)
			trimmed = strings.TrimRight(trimmed, "\n")
			trimmed = strings.TrimLeft(trimmed, "\n")
			if len(trimmed) > 0 {
				depth++
				fmt.Printf("\n%*s%s", depth*2, "", n.Data)
				depth--
			}
		case html.ErrorNode:
		}
	}
}
func printDoctypeElement(n *html.Node) {
	fmt.Printf("<!DOCTYPE %s>\n", n.Data)
}

func endElement(n *html.Node) {
	if n.Type != html.ElementNode {
		return
	}
	depth--
	if isAvailableChildNode(n.FirstChild) {
		fmt.Printf("\n%*s</%s>", depth*2, "", n.Data)
	}
}

func printStartElement(n *html.Node) {
	if isAvailableChildNode(n.FirstChild) {
		fmt.Printf("\n%*s<%s%s>", depth*2, "", n.Data, appendAttrs(n))
	} else if n.Data == "br"{
		fmt.Printf("<%s/>", n.Data)
	} else {
		fmt.Printf("\n%*s<%s%s />", depth*2, "", n.Data, appendAttrs(n))
	}
	depth++
}

func appendAttrs(n *html.Node) (attrs string) {
	if len(n.Attr) > 0 {
		for _, attr := range n.Attr {
			attrs += fmt.Sprintf(" %s=%#v", attr.Key, attr.Val)
		}
	}
	return attrs
}

func isAvailableChildNode(n *html.Node) bool {
	if n == nil {
		return false
	}
	return true
}

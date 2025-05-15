package main

// way to run this code
// fetch https://golang.org | findlinks1

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

func main(){
	doc, err := html.Parse(os.Stdin)
	if err!= nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
    for _, link := range visit(nil, doc){
		fmt.Println(link)
	}
}

// this fetches the links

// func visit(links []string, n *html.Node) []string{
// 	if n.Type == html.ElementNode && n.Data =="a"{
// 		for _, a := range n.Attr{
// 			if a.Key == "href" {
// 				links = append(links, a.Val)
// 			}
// 		}
// 	}
// 	for c:= n.FirstChild; c!= nil; c = c.NextSibling{
// 		links = visit(links, c)
// 	}
// 	return links
// }

// this fetches the image sources

func visit(links []string, n *html.Node) []string{
	if n.Type == html.ElementNode && n.Data =="img"{
		for _, a := range n.Attr{
			if a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	for c:= n.FirstChild; c!= nil; c = c.NextSibling{
		links = visit(links, c)
	}
	return links
}
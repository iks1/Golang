package main

// this code merges ch1/fetch.go and ch5/findlinks1.go
// introducing error control

import (
	"os"
	"fmt"
	"golang.org/x/net/html"
	//"io/ioutil"
	"net/http"
)

func main(){
	for _, url := range os.Args[1:]{
		//fmt.Println(url)
		links, err := findLinks(url)
		if err!= nil{
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findLinks(url string) ([]string, error){
	resp, err := http.Get(url)
	if err!= nil {
		return nil, err
	}
    if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s : %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err!= nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

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
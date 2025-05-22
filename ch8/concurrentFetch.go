package main

/*
  running each iteration of a loop as a separate goroutine

  wg.Add(1) tells we are adding one more go routine 
  wg.Done() decrements the counter by 1
  wg.Wait() Blocks the main goroutine until the counter becomes 0
*/

import (
	"fmt"
	"net/http"
	"sync"
)

func main(){
	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://example.com",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1) 
		go func(u string){
			defer wg.Done()
			resp, err := http.Get(u)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(u, "->", resp.Status)
			resp.Body.Close()
		}(url)
	}

	wg.Wait() // wait for all the goroutines to finish
}

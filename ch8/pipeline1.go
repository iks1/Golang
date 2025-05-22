package main

// this program connects three go routines 
// making a pipeline

import(
	"fmt"
)

func main(){
	naturals := make(chan int)
	squares := make(chan int)

	//counter 
	go func(){
		for x:=0; ;x++{
			naturals <- x}
		}()

		go func(){
			for {
				x:= <- naturals
				squares <- x*x
			}
		}()

		for {
			fmt.Println(<-squares)
		}
	}
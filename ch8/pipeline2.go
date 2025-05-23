package main

// pipeline
// you can close the channles when 
// it is needed to tell the receiver that 
// no more data will be sent
import(
	"fmt"
)

func main(){
	naturals := make(chan int)
	squares := make(chan int)

	//counter 
	go func(){
		for x:=0; x<100 ;x++{
			naturals <- x}
			close(naturals)
		}()

		go func(){
			for {
				x:= <- naturals
				squares <- x*x
			}
			close(squares)
		}()

		for x := range squares {
			fmt.Println(x)
		}
	}
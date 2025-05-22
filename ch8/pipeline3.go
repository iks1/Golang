package main
import "fmt"

/*
 unidirectional channels 

 chan <- send only
 <- receive only

 restricts what the function should do
*/

func main(){
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)

	go squares(naturals, squares)

	print(squares)
}

// this function only sends data (allowed to)
func counter(out chan <- int){
	for x:=0; ; x++ {
		out <-x
	}
}


// this function does both (allowed to)
func squares(in <-chan int, out chan <- int){
	for{
		x := <- in 
		out <- x*x
	}
}


// this function only takes data (allowed to)
func printer(in <-chan int){
	for{
		fmt.Println(<-in)
	}
}
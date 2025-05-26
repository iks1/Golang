package main 
import(
	"fmt"
	"sync"
)

var once sync.once 

//gets called only once 
func initialize(){
	fmt.Println("starting")
}

func main(){
	for i:=0; i<5; i++{
		go func(){
			once.Do(initialize)
		}()
	}
	//waiting for goroutines to finish
	select{}
}

//loadIcons example from the book
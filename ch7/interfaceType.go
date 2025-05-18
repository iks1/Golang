package main

import(
	"fmt"
)

func main(){
	mysterBox := interface{}(10)
	describeValue(mysterBox)

	retrivedInt, ok = mysterBox.(int)

	if ok{
		fmt.Println("Retrieved int: ", retrivedInt)
	}
	else {
		fmt.Println("Value is not an integer")
	}
}

func describeValue(t interface{}){
	fmt.Printf("Type: %T, Value: %v\n",t ,t)
}
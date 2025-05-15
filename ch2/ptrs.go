package main

import (
	"fmt"
)

//pointers work same as c/c++
func main(){
	var x int
	x := 1  // this is declaration with initialization and this will throw
	        // error as already declared abhove instead use x=1 or remove the
			// abhove line
	p := &x // p is type of *int, pointing to x
	fmt.Println(*p)
	fmt.Println(x)
    *p =2   // same as x=2
	fmt.Println(x)
}

/*
var p = f()

func f() *int{
	v:=1
	return &v
}
*/

/*
func incr(p *int) int {
   *p++
   return *p
}

v:=1
incr(&v)
fmt.Println(incr(&v)) //3
*/

// new(T) returns pointer to variable
/*
p := new(int) // p, of type *int
fmt.Println(*p)
*p=2
fmt.Println(*p)
*/




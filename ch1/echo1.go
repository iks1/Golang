package main

import (
	"fmt"
	"os"
	"strings"
)

// the following programs give same output

// func main(){
// 	var s, sep string
// 	for i:= 1; i<len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main(){
// 	s, sep := "", ""
// 	for _, arg := range os.Args[1:]{
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

func main(){
   fmt.Println(strings.Join(os.Args[0:], " "))
   //fmt.Println(os.Args[1:]
}

/*
All are equivalent

s := ""
var s string
var s = ""
var s string = ""

*/
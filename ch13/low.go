package main 

import (
	_ "os"
	"fmt"
	_ "net/http"
	"unsafe"
)

func main(){
    fmt.Println(unsafe.Sizeof(float64(0)))
}
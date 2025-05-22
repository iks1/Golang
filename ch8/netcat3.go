package main

// struct{}{} is zero size struct, preferred when you do
// not need to send any data 

// make the program wait for the background goroutine 
// to complete before exiting

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func(){
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} //signal the main goroutine
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<- done //wait to backgrount to finish
}
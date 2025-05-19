package main

import(
	"fmt"
	"os"
	"io"
)

//implementing something like Fprintf/ Sprintf/ Printf


//package io
type Writer interface{
   Write(p []byte)(n int, err error)
}


//inside package fmt

func Fprintf(w Writer, format string, args ...interface{}) (int, error)
func Printf(format string, args ...interface{}) (int, error){
	return Fprintf(os.Stdout, format, args...)
}
func Sprintf(w Writer, format string, args ...interface{}) (int, error){
	var buf bytes.Buffer
	Fprintf(&buf, format, args...)
	return buf.string()
}

func main(){

}
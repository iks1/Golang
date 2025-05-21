package main

/*
type Value interface{
  String() string 
  Set(string) error
}
*/

import(
	"fmt"
	"flag"
	"strconv"
)

type Celsius float64

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error{
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	f.Celsius = Celsius(value)
	return nil 
}

func (f *celsiusFlag) string() string {
	return fmt.Sprintf(".1fC", f.Celsius)
}

func CelsiusFlag(name string, defaultVal Celcius, usage string) *Celsius {
	f := celsiusFlag{defaultVal}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celcius
}

func main(){
    temp := CelsiusFlag("temp", 20.0, "the temparature in Celcius")
	flag.Parse()
	fmt.Println(*temp)
}
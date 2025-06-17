package main

const prefix = "Hello, "

func Helloji(name string) string {
	if name == "" {
		name = "world"
	}
	return prefix + name
}

// func main(){
// 	fmt.Println(Hello())
// }
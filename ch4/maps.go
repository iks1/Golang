ages := make(map[string]int)
ages["alice"] = 31
ages["bob"] = 30

//same as 
ages := map[string]int{
	"alice": 31,
	"bod":30,
}

//accessing 
for name, age := range ages{
	fmt.Printf("%s\t%d\n", name, age)
}
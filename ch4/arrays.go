var a [3]int // array of 3 integers 
fmt.Println(a[0])

//printing the array
for i, v := range a{
	fmt.Printf("%d %d", i, v)
}

//ignoring index
for _, v := range a{
	fmt.Println("%d\n", v)
}

var a [3]int = [3]{1,2,3}
// or can be initialized as 
a := [...]int{1,2,3}


type Currency int

const (
	USD Currency = iota //0 
	EUR //1
	GBP //2
	RMB //3
)
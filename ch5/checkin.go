//variadic functions

func sum(vals ...int) int {
	total :=0
	for _,val := range vals{
		total += val
	}
	return total
}

sum(2,3,4,5)

vals := []int{1,2,3,4}
sum(vals...)


//defer functions

func lookup(key string) int {
	mu.lock()
	defer mu.unlock()
	return m[key]
}
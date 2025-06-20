package main 

func sum(numbers []int) int {
	ans := 0

	for _, val := range numbers {
		ans = ans+val
	}
	return ans
}

//variadic
func sumAll(numbersToSum ...[]int) []int{
	var ans []int 
	for _, numbers := range numbersToSum{
		ans = append(ans, sum(numbers))
	}
	return ans
}

func getTailSum(numsToGetTail ...[]int) []int {
	var ans []int 
	for _, numbers := range numsToGetTail{
		if len(numbers) <= 1 {
		   ans = append(ans, 0)	
		} else {
			ans = append(ans, sum(numbers[1:]))
		}
	}
	return ans
}

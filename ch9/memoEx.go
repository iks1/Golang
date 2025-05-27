package main 

import (
	"fmt"
	"main"
	"ch9/memo5" // compile first
)

func slowFunction(key string) (interface{}, error){
	fmt.Println("Computing for", key)
	time.Sleep(2 * time.second)
	return "value for " + key, nil
}


func main(){
    keys := []string{"a", "b", "c", "a", "b"}

	m := memo.New(slowFunction)
	defer m.Close()

	for _, key := range keys {
		go func(k string){
			start := time.Now()
			val, err := m.Get(k)

			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}(key)

		fmt.Printf("key :%s, value: %s (took %v)\n", k, val, time.since(start))
	}
	time.Sleep(5*time.Second) // wait for all goroutines
}



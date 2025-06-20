package main 

import (
	"testing"
    "reflect")

func TestSum(t *testing.T){
 	t.Run("testing sum function",func(t *testing.T){
        numbers := [5]int{1,2,3,4,5}
		got := sum(numbers[:])
		want := 15

		if got!= want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})

	t.Run("testing sumAll function", func(t *testing.T){
		got := sumAll([]int{1,2}, []int{0,9})
		want := []int{3,9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v ", got, want)
		}
	})
}

func TestTailSums(t *testing.T){
	t.Run("normal test", func(t *testing.T){
       got := getTailSum([]int{3,4,5}, []int{6,7,7})
	   want := []int{9,14}
	   if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v ", got, want)
	    }
	})

	t.Run("base case test", func(t *testing.T){
	   got := getTailSum([]int{}, []int{9})
	   want := []int{0,0}
	   if !reflect.DeepEqual(got, want){
		t.Errorf("got %v want %v", got, want)
	   }

	})
}
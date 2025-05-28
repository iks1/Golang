package word 
import "testing"

func TestPalindrome(t *testing.T){
	if !IsPalindrome("detartrated"){
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T){
	if IsPalindrome("palindrome"){
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}
package main

import "testing"

func TestHelloji(t *testing.T){
	t.Run("saying hello to people", func(t *testing.T){
		got := Helloji("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'hello world' default", func(t *testing.T){
		got := Helloji("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got!= want{
		t.Errorf("got %q want %q", got, want)
	}
}
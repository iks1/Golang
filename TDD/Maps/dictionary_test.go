package main 

import "testing"

func TestSearch(t *testing.T){
	dictionary := Dictionary{"test":"this is just a test"}

	t.Run("known word", func(t *testing.T){
       got, _ := dictionary.Search("test")
	   want := "this is just a test"
       assertStrings(t, got, want)

	})

    t.Run("known word", func(t *testing.T){
		_, err := dictionary.Search("unknown")
		want := ErrorNotFound

        if err == nil {
			t.Fatal("expected to get an error")
		}

		assertErrors(t, err, want)
	})
}

func TestAdd(t *testing.T){
	t.Run("new word", func(t *testing.T){
        dictionary := Dictionary{}
		word := "test"
		defination := "this is just a test"
		err := dictionary.Add(word, defination)
        
		assertErrors(t, err, nil)
		assertDefination(t, dictionary, word, defination)
	})

	t.Run("existing word", func(t *testing.T){
		word := "test"
		defination := "this is just a test"
		dictionary := Dictionary{word : defination}
		err := dictionary.Add(word, "new test")

		assertErrors(t, err, ErrWordExists)
		assertDefination(t, dictionary, word, defination)
	})
}

func TestUpdate(t *testing.T){
	t.Run("existing word", func(t *testing.T){
      	word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new Definition"

		err := dictionary.Update(word, newDefinition)
		assertErrors(t, err, nil)
		assertDefination(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T){
		word:= "test" 
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T){
	t.Run("word exist", func(t *testing.T){
        word := "test"
		dictionary := Dictionary{word: "test definition"}
		err := dictionary.Delete(word)

		assertErrors(t, err, nil)

		_, err = dictionary.Search(word)
		assertErrors(t, err, ErrorNotFound)
	})

	t.Run("word does not exist", func(t *testing.T){
        word := "test"
		dictionary := Dictionary{}
		err:= dictionary.Delete(word)
		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertErrors(t testing.TB, got, want error){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefination(t *testing.T, dictionary Dictionary, word, defination string){
    t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word", err)
	}
	assertStrings(t, got, defination)
}

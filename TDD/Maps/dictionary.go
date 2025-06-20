package main 

import (
	"errors"
)

type Dictionary map[string]string

var (
	ErrWordExists = errors.New("could not word because it already exists")
	ErrorNotFound = errors.New("could not find the word you were looking for")
)

type DictionaryErr string 

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	defination, ok := d[word]

	if !ok {
		return "", ErrorNotFound
	}

	return defination, nil
}

func (d Dictionary) Add(word, defination string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = defination
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string){
	d[word] = definition
}

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }


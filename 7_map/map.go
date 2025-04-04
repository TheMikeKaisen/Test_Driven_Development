package dictionary

import "errors"

var ErrNotFound error = errors.New("could not find the word you were searching for")

func Search(dict map[string]string, word string) (string, error) {
		if dict[word] == "" {
		return "", errors.New("no such sentence exist")
	}
	return dict[word], nil; 
}

type Dictionary map[string]string
func (d Dictionary) Search(word string) (string, error) {
	if (d[word] == "") {
		return "",errors.New(ErrNotFound.Error())
	}
	return d[word], nil
}



func (d Dictionary) Add(word, value string) {
	d[word] = value
}
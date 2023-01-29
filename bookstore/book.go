package bookstore

import "errors"

type Book struct {
	// ...
	Copies int
	Title  string
	Author string
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies -= 1
	return b, nil
}

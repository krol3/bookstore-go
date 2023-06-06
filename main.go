package main

import (
	"fmt"

	"github.com/krol3/bookstore-go/bookstore"
)

func main() {

	books := []bookstore.Book{
		{Title: "Delightfully Uneventful Trip on the Orient Express"},
		{Title: "One Hundred Years of Good Company"},
	}

	fmt.Printf("Books length: %s \n", len(books))

	b := bookstore.Book{Title: "The Grapes of Mild Irritation"}
	books = append(books, b)

	fmt.Println(books[2].Title)
}

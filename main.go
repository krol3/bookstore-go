package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
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

	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Printf("Some error occured. Err: %s", err)
	}

	val := os.Getenv("AWS_SECRET_ACCESS_KEY")
	fmt.Println(val)
}

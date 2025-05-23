package main

type Book struct {
	Title  string
	Author string
	Year   int
	Genre  string
}

func NewBook(title string, author string, year int, genre string) *Book {
	return &Book{title, author, year, genre}
}

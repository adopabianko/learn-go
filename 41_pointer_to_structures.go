package main

import "fmt"

type Books struct {
	Title   string
	Author  string
	Subject string
	BookID  int
}

func main() {
	var Book1 Books /* Declare Book1 of type Book */
	var Book2 Books /* Declare Book2 of type Book */

	/* book 1 specification */
	Book1.Title = "Go Programming"
	Book1.Author = "Mahesh Kumar"
	Book1.Subject = "Go Programming Tutorial"
	Book1.BookID = 6495407

	/* book 2 specification */
	Book2.Title = "Telecom Billing"
	Book2.Author = "Zara Ali"
	Book2.Subject = "Telecom Billing Tutorial"
	Book2.BookID = 6495700

	/* print Book1 info */
	printBook(&Book1)

	/* print Book2 info */
	printBook(&Book2)
}
func printBook(book *Books) {
	fmt.Printf("Book Title : %s\n", book.Title)
	fmt.Printf("Book Author : %s\n", book.Author)
	fmt.Printf("Book Subject : %s\n", book.Subject)
	fmt.Printf("Book BookID : %d\n", book.BookID)
}

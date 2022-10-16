//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

const (
	CHECKED_IN  = true
	CHECKED_OUT = false
)

type BookStatus bool

type Book struct {
	name           string
	checkOutStatus BookStatus
	inTime         time.Time
	outTime        time.Time
}

type Member struct {
	name string
}

type Library struct {
	members         []Member
	books           []Book
	booksCheckedOut []Book
}

func createBooks() []Book {
	books := []Book{
		{"Java", CHECKED_IN, time.Now(), time.Time{}},
		{"C++", CHECKED_IN, time.Now(), time.Time{}},
		{"Python", CHECKED_OUT, time.Time{}, time.Now()},
		{"Angular", CHECKED_OUT, time.Time{}, time.Now()},
	}
	return books
}

func createMembers() []Member {
	member := []Member{
		{"Naruto"},
		{"Goku"},
		{"Luffy"},
	}
	return member
}

func main() {
	library := Library{
		createMembers(),
		createBooks(),
		nil,
	}
	fmt.Println(library)

	checkOut(&library.books[1])
	library.booksCheckedOut = append(library.booksCheckedOut, library.books[1])
	fmt.Println(library)

	checkIn(&library.books[3])
	fmt.Println(library)

}

func checkOut(book *Book) {
	book.outTime = time.Now()
	book.checkOutStatus = CHECKED_OUT
}

func checkIn(book *Book) {
	book.inTime = time.Now()
	book.checkOutStatus = CHECKED_IN
}

package main

import "fmt"

func main() {
	// embedMethods()
	embedInterface()
}

type Author struct {
	AuthorName string
	AuthorAge  int
}

func (author Author) SayHello() string {
	return fmt.Sprintf("I am %v. Hello GoLinuxCloud members!", author.AuthorName)
}

type Book struct {
	Title string
	Author
}

// The book struct above will equivalent to:
// type Book struct {
//    Title string
//    AuthorName string
//    AuthorAge int
// }

func embedMethods() {
	book1 := Book{
		Title: "Book1",
		Author: Author{
			AuthorName: "author1",
			AuthorAge:  19,
		},
	}

	author2 := Author{
		AuthorName: "author2",
		AuthorAge:  19,
	}

	fmt.Println(book1)
	fmt.Println(book1.AuthorAge)
	fmt.Println(book1.Author.AuthorAge)

	fmt.Println(author2.SayHello())
	fmt.Println(book1.SayHello())
}

type interface1 interface {
	AuthorDetails()
}

type interface2 interface {
	CustomDetails()
}

type embedded interface {
	interface1
	interface2
}

func (a Author) AuthorDetails() {
	fmt.Printf("I am %v, I am %v years old. Hello GoLinuxCloud members!\n", a.AuthorName, a.AuthorAge)
}

func (a Author) CustomDetails() {
	fmt.Printf("This is a customized message! I'm %v, I am %v years old.\n", a.AuthorName, a.AuthorAge)
}

func embedInterface() {
	author := Author{
		AuthorName: "author1",
		AuthorAge:  22,
	}

	// Accessing the methods of the interface1 and interface2
	var f embedded = author
	f.AuthorDetails()
	f.CustomDetails()
}

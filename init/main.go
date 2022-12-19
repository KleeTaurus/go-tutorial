package main

import (
	"fmt"

	"github.com/KleeTaurus/go-tutorial/init/a"
	"github.com/KleeTaurus/go-tutorial/init/b"
)

var greetings string
var age int

func init() {
	fmt.Println("I always execute before main() function")
	greetings = "Hello world"
}

func init() {
	fmt.Println("<<< First >>>")
}

func init() {
	fmt.Println("<<< Second >>>")
}

func init() {
	fmt.Println("<<< Third >>>")
}

func main() {
	multiPackages()
}

func simpleInit() {
	fmt.Println("I execute after init() function")
	fmt.Println(greetings)
	fmt.Printf("Go language is %d years old\n", age)
}

func multiPackages() {
	fmt.Println("Executing main() function in main.go")
	b.Greetings()
	a.Greetings()
}

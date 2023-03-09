package main

import "fmt"

type Options struct {
	Name      string
	Age       int
	IsStudent bool
}

type Person struct {
	name      string
	age       int
	isStudent bool
}

type Option func(*Options)

func WithName(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

func WithAge(age int) Option {
	return func(o *Options) {
		o.Age = age
	}
}

func WithIsStudent(isStudent bool) Option {
	return func(o *Options) {
		o.IsStudent = isStudent
	}
}

func NewPerson(opts ...Option) *Person {
	options := &Options{
		Name:      "Unknown",
		Age:       18,
		IsStudent: false,
	}

	for _, opt := range opts {
		opt(options)
	}

	return &Person{
		name:      options.Name,
		age:       options.Age,
		isStudent: options.IsStudent,
	}
}

func main() {
	p1 := NewPerson()
	fmt.Println(p1)

	p2 := NewPerson(WithName("Bob"), WithAge(20), WithIsStudent(true))
	fmt.Println(p2)
}

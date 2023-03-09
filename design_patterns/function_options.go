package main

import "fmt"

// 在上面的代码中，我们定义了一个 Options 结构体类型，该结构体包含了所有可选
// 参数。我们还定义了一个 Option 类型，该类型是一个函数类型，用于设置 Options
// 结构体中的可选参数。
//
// 接下来，我们定义了一些函数，例如 WithName、WithAge 和 WithIsStudent，这些
// 函数都是 Option 类型的函数，用于设置 Options 结构体中的对应可选参数。
//
// 最后，我们定义了一个 NewPerson 函数，该函数接受多个 Option 类型的参数，用
// 于设置 Options 结构体中的所有可选参数。我们根据 Options 结构体中的可选参数
// 创建了一个 Person 对象，并将其返回。
//
// 在 main 函数中，我们首先使用 NewPerson 函数创建了一个默认的 Person 对象，然
// 后使用 WithName、WithAge 和 WithIsStudent 函数自定义了一个 Person 对象，并
// 输出了这两个对象的属性值。
//
// 通过使用 Function Options Pattern，我们可以在 Golang 中更加灵活地定义方法的
// 可选参数，并且使得代码更加易读易用。

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

package main

import (
	"fmt"
	"time"
)

func main() {
	// timeAndSince()
	// timeAndSub()
	withDefer()
}

func timeAndSince() {
	start := time.Now()

	// sum of 10000 consecutive numbers
	sum := 0
	for i := 1; i < 10000; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)

	// calculate to exe time
	elapsed := time.Since(start)
	fmt.Printf("Sum function took %s\n", elapsed)
}

func timeAndSub() {
	start := time.Now()

	// sum of 10000 consecutive numbers
	sum := 0
	for i := 1; i < 10000; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)

	// calculate to exe time
	// using time.Sub() function
	elapsed := time.Now().Sub(start)
	fmt.Printf("Sum function took %s\n", elapsed)
}

func withDefer() {
	defer exeTime("main")()

	// sum of 10000 consecutive numbers
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
		// sleep 1s
		time.Sleep(time.Second * 1)
	}
	fmt.Println("sum:", sum)
}

// func to calculate and print execution time
func exeTime(name string) func() {
	fmt.Println("exeTime function was called")
	start := time.Now()
	return func() {
		fmt.Println("anonymous function was called")
		fmt.Printf("%s execution time: %v\n", name, time.Since(start))
	}
}

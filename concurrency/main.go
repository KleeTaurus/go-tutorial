package main

import (
	"fmt"
	"time"
)

func main() {
	// simpleGoroutine()
	// complexGoroutine()
	// makeChannel()
	sendAndReceive()
}

// Prints numbers from 1-3 along with the passed string
func foo(s string) {
	for i := 1; i <= 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, ": ", i)
	}
}

func simpleGoroutine() {
	fmt.Println("simpleGoroutine started...")
	// Starting two goroutines
	go foo("1st goroutine")
	go foo("2nd goroutine")

	// Wait for goroutines to finished before main goroutine ends
	time.Sleep(time.Second)
	fmt.Println("simpleGoroutine finished")
}

func complexGoroutine() {
	messages := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		messages <- 1
	}()
	go func() {
		time.Sleep(time.Second * 2)
		messages <- 2
	}()
	go func() {
		time.Sleep(time.Second * 1)
		messages <- 3
	}()
	go func() {
		for i := range messages {
			fmt.Println(i)
		}
	}()
	go func() {
		time.Sleep(time.Second * 1)
		messages <- 4
	}()
	go func() {
		time.Sleep(time.Second * 1)
		messages <- 5
	}()
	time.Sleep(time.Second * 5)
}

func makeChannel() {
	var c chan int
	if c == nil {
		fmt.Println("channel c is nil, going to define it")
		c = make(chan int)
		fmt.Printf("Type of c is %T\n", c)
	}
}

func sendValues(myIntChannel chan int) {
	for i := 0; i < 5; i++ {
		myIntChannel <- i
	}

	close(myIntChannel)
}

func sendAndReceive() {
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for i := 0; i < 6; i++ {
		val, ok := <-myIntChannel
		fmt.Println("val: ", val, ", ok: ", ok)
	}
}

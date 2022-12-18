package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	// shouldPanic() // the variable employees should be initialized with make(...)
	// initWithMapLiteral()
	// initWithMake()
	// addingKVPair()
	// accessingValues()
	// changingValue()
	// deletingElement()
	// ifKeyExists()
	// forRangeLoop()
	forRangeLoopThroughOrderedMap()
}

func shouldPanic() {
	var employees map[string]int
	fmt.Println(employees)
	fmt.Println(reflect.TypeOf(employees))

	if employees == nil {
		fmt.Println("Employee map is nil")
	} else {
		fmt.Println("Employee map is not nil")
	}

	employees["John Doe"] = 1
	fmt.Println(employees)
}

func initWithMapLiteral() {
	employees := map[string]int{
		"John Doe": 23,
		"Mary Doe": 19,
	}
	fmt.Println(employees)

	if employees == nil {
		fmt.Println("Employee map is nil")
	} else {
		fmt.Println("Employee map is not nil")
	}
}

func initWithMake() {
	employees := make(map[string]int, 5)
	fmt.Println(employees)

	if employees == nil {
		fmt.Println("Employee map is nil")
	} else {
		fmt.Println("Employee map is not nil")
	}

	employees["John Doe"] = 23
	employees["Mary Doe"] = 19

	fmt.Println(employees)
}

func addingKVPair() {
	employees := make(map[string]int)
	employees["John Doe"] = 45
	employees["Stephanie Doe"] = 21
	employees["Tom Doe"] = 33
	employees["Mary Doe"] = 12

	fmt.Println(employees)
}

func accessingValues() {
	employees := make(map[string]int)
	employees["John Doe"] = 45
	employees["Stephanie Doe"] = 21
	employees["Tome Doe"] = 33
	employees["Mary Doe"] = 12

	employeeOne := employees["John Doe"]
	newEmployee := employees["Mike"]

	fmt.Println(employeeOne)
	fmt.Println(newEmployee)
}

func changingValue() {
	employees := make(map[string]int)
	employees["John Doe"] = 45
	employees["Stephanie Doe"] = 21
	employees["John Doe"] = 40
	employees["Stephanie Doe"] = 30
	fmt.Printf("Updated employees map %v\n", employees)
}

func deletingElement() {
	employees := make(map[string]int)
	employees["John Doe"] = 45
	employees["Stephanie Doe"] = 21

	delete(employees, "John Doe")
	fmt.Println(employees)
}

func ifKeyExists() {
	users := map[string]int{
		"John":  21,
		"David": 43,
		"Paul":  54,
	}

	value, exist := users["John"]
	if exist {
		fmt.Printf("Found %d %v\n", value, exist)
	} else {
		fmt.Printf("Not found %d %v\n", value, exist)
	}
}

func forRangeLoop() {
	users := map[string]int{
		"John":  21,
		"David": 12,
		"Paul":  23,
	}

	for userName, userAge := range users {
		fmt.Printf("%s is %d years old\n", userName, userAge)
	}
}

func forRangeLoopThroughOrderedMap() {
	users := map[int]string{
		100: "John",
		23:  "Mary",
		54:  "Paul",
		21:  "Josh",
	}
	var keys []int

	for key := range users {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		fmt.Printf("Key: %d and Value: %s\n", key, users[key])
	}
}

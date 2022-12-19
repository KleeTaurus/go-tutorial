package main

import "fmt"

func main() {
	// typeAssertion()
	// okCommaIdiom()
	typeSwitch()
}

func typeAssertion() {
	// an interface which has a string value
	var checkInterface interface{} = "GoLinuxCloud"

	// assigning value of interface type to checkType variable
	checkType := checkInterface.(string)

	// printing the concrete value
	fmt.Println(checkType)

	// panic because interface does not have int type
	checkTypeInt := checkInterface.(int)

	fmt.Println(checkTypeInt)
}

func okCommaIdiom() {
	// an interface has a float64 value
	var checkInterface interface{} = 11.04

	// assigning value of interface type to checkType variable
	checkType := checkInterface.(float64)

	// printing the concrete value
	fmt.Println(checkType)

	// test whether interface has string type and
	// return true if found or false otherwise
	checkType2, ok := checkInterface.(string)
	if ok {
		fmt.Println("Correct assertion!")
		fmt.Println(checkType2)
	} else {
		fmt.Printf("Wrong assertion, got data of type %T but wanted String!", checkInterface)
	}
}

func typeSwitch() {
	var testInterface interface{} = "GoLinuxCloud"
	var testInterface2 interface{} = map[string]int{
		"Anna":  4,
		"Bob":   10,
		"Clair": 11,
	}

	var checkType = func(i interface{}) {
		switch i.(type) {
		case string:
			fmt.Println(i, "is a string!")
		case int:
			fmt.Println(i, "is a int!")
		case float64:
			fmt.Println(i, "is a float64!")
		default:
			fmt.Println(i, "is not basic type!")
		}
	}

	checkType(testInterface)
	checkType(testInterface2)
}

package main

import (
	"fmt"
	"os"
)

type salaryCalculator interface {
	calculateSalary() float64
	report()
}

type PermanentEmployee struct {
	id          int
	basicSalary float64
	commission  float64
}

type ContractEmployee struct {
	id          int
	basicSalary float64
}

func (p PermanentEmployee) calculateSalary() float64 {
	return p.basicSalary + (p.commission/100)*p.basicSalary
}

func (c ContractEmployee) calculateSalary() float64 {
	return c.basicSalary
}

func (p PermanentEmployee) report() {
	fmt.Printf("Employee ID %d earns USD %f per month\n", p.id, p.calculateSalary())
}

func (c ContractEmployee) report() {
	fmt.Printf("Employee ID %d earns USD %f per month\n", c.id, c.calculateSalary())
}

type ReadFileError struct {
	message string
}

func (rfe ReadFileError) Error() string {
	return fmt.Sprintf("Custom Read File Error! Error Message: %s", rfe.message)
}

func errorHandler(err error) {
	fmt.Println(err.Error())
}

func main() {
	// emptyInterface()
	// implementInterface()
	// calculateTotalSalary()
	// interfaceAsArgument()
	errorInterface()
}

func emptyInterface() {
	var s salaryCalculator
	fmt.Println(s)
}

func implementInterface() {
	var calculator salaryCalculator
	calculator = PermanentEmployee{id: 1, basicSalary: 10000, commission: 20}
	calculator.report()
	calculator = ContractEmployee{id: 2, basicSalary: 5000}
	calculator.report()
}

func calculateTotalSalary() {
	p1 := PermanentEmployee{id: 1, basicSalary: 2300, commission: 13}
	p2 := PermanentEmployee{id: 2, basicSalary: 1500, commission: 18}
	p3 := PermanentEmployee{id: 3, basicSalary: 2300, commission: 10}
	c1 := ContractEmployee{id: 4, basicSalary: 500}
	c2 := ContractEmployee{id: 5, basicSalary: 1100}
	c3 := ContractEmployee{id: 6, basicSalary: 700}

	employees := []salaryCalculator{p1, p2, p3, c1, c2, c3}

	var totalSalary float64
	for _, employee := range employees {
		totalSalary += employee.calculateSalary()
	}
	fmt.Printf("Company total salary is USD %f\n", totalSalary)
}

func salaryExpense(employees []salaryCalculator) float64 {
	var totalSalary float64
	for _, employee := range employees {
		totalSalary += employee.calculateSalary()
	}

	return totalSalary
}

func interfaceAsArgument() {
	p1 := PermanentEmployee{id: 1, basicSalary: 2300, commission: 13}
	p2 := PermanentEmployee{id: 2, basicSalary: 1500, commission: 18}
	p3 := PermanentEmployee{id: 3, basicSalary: 2300, commission: 10}
	c1 := ContractEmployee{id: 4, basicSalary: 500}
	c2 := ContractEmployee{id: 5, basicSalary: 1100}
	c3 := ContractEmployee{id: 6, basicSalary: 700}

	employees := []salaryCalculator{p1, p2, p3, c1, c2, c3}
	totalSalary := salaryExpense(employees)
	fmt.Printf("Company total salary is USD %f\n", totalSalary)
}

func errorInterface() {
	_, err := os.ReadFile("logs.txt")
	if err != nil {
		e := ReadFileError{message: "Error reading file"}
		errorHandler(e)
		return
	}
	fmt.Println("Read data file file")
}

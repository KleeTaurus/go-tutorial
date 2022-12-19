package main

import "fmt"

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	// simpleGeneric()
	// noneGeneric()
	// withGeneric()
	typeConstraint()
}

func simpleGeneric() {
	ints := []int{1, 2, 3}
	strings := []string{"One", "Two", "Three"}
	Print(ints)
	Print(strings)
}

func SumOfFloat(nums []float64) float64 {
	var sum float64

	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumOfIntegers(nums []int64) int64 {
	var sum int64

	for _, num := range nums {
		sum += num
	}

	return sum
}

func noneGeneric() {
	f := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	i := []int64{1, 2, 3, 4, 5}

	s1 := SumOfFloat(f)
	s2 := SumOfIntegers(i)

	fmt.Println("Sum of float64: ", s1)
	fmt.Println("Sum of int64: ", s2)
}

func genericSum[N int64 | float64](nums []N) N {
	var sum N

	for _, num := range nums {
		sum += num
	}

	return sum
}

func withGeneric() {
	f := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	i := []int64{1, 2, 3, 4, 5}

	s1 := genericSum(f)
	s2 := genericSum(i)

	fmt.Println("Sum of float64: ", s1)
	fmt.Println("Sum of int64: ", s2)
}

type Number interface {
	float64 | int64
}

func genericSumV2[N Number](nums []N) N {
	var sum N

	for _, num := range nums {
		sum += num
	}

	return sum
}

func typeConstraint() {
	f := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	i := []int64{1, 2, 3, 4, 5}

	s1 := genericSumV2(f)
	s2 := genericSumV2(i)

	fmt.Println("Sum of float64: ", s1)
	fmt.Println("Sum of int64: ", s2)
}

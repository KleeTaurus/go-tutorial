package main

import (
	"fmt"
	"reflect"
)

func main() {
	// nothingMagical()
	sqlInsertQuery()
	// extractValue()
}

func nothingMagical() {
	i := 10
	fmt.Printf("%d %T", i, i)
}

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func sqlInsertQuery() {
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("Unsupported type")
}

func extractValue() {
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	fmt.Println("Name: ", reflect.TypeOf(a).Name())

	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)
}

package main

import (
	"fmt"
	"log"
	"math"

	intslice "github.com/KleeTaurus/go-tutorial/packages/slice/intslice"
	str "github.com/KleeTaurus/go-tutorial/packages/string"
	"gopkg.in/yaml.v2"
)

func main() {
	// exportedFunction()
	// accessPackage()
	encodeAndDecodeYAML()
}

func exportedFunction() {
	// exported Min function
	fmt.Println(math.Min(10, 91))

	// exported Sqrt function
	fmt.Println(math.Sqrt(144))

	// the value of `𝜋`
	fmt.Println(math.Pi)
}

func accessPackage() {
	fmt.Println(intslice.SortInt([]int{14, 5, 19, 10}))
	fmt.Println(str.Upper("golinuxcloud"))
}

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func encodeAndDecodeYAML() {
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}

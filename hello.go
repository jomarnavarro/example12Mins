package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	helloWorld()
	vars()
	arrays()
	arrays2()
	slices()
	maps()
	loops()
	loops2()
	iterateSlices()
	iterateMaps()
	fmt.Println(sum(1, 2))
	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	result, err = sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println("types")
	p := person{name: "Omar", age: 23, salary: 23000.00}
	p2 := person{age: 24, salary: 25000.00, name: "Monica"}
	p3 := person{age: 14, name: "Nena"}
	fmt.Println(p, " ", p2, " ", p3)
	fmt.Println("only a field, name: ", p.name)

	pointers()
}

func helloWorld() {
	fmt.Println("Hello world!")
}

func vars() {
	fmt.Println("vars")
	var x int = 5

	y := 6
	fmt.Println(x + y)
}

func arrays() {
	fmt.Println("arrays")
	var a [5]int
	a[2] = 7
	fmt.Println(a)
}

func arrays2() {
	fmt.Println("arrays 2")
	a := [5]int{5, 4, 3, 2, 1}

	fmt.Println(a)
}

func slices() {
	fmt.Println("slices")
	a := []int{5, 4, 3}
	a = append(a, 12)

	fmt.Println(a)
}

func maps() {
	fmt.Println("maps")
	vertix := make(map[string]int)

	vertix["triangle"] = 2
	vertix["square"] = 3
	vertix["dodecagon"] = 12
	delete(vertix, "square")
	fmt.Println(vertix)
}

func loops() {
	fmt.Println("for loops")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func loops2() {
	fmt.Println("whileish loop")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func iterateSlices() {
	fmt.Println("iterate slices")
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index: ", index, " value: ", value)
	}
}

func iterateMaps() {
	fmt.Println("iterate maps")
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"
	m["c"] = "gama"
	m["d"] = "delta"

	for key, value := range m {
		fmt.Println("key: ", key, " value: ", value)
	}
}

func sum(x int, y int) int {
	fmt.Println("doing a sum")
	return x + y
}

func sqrt(x float64) (float64, error) {
	fmt.Println("getting a squareRoot")
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

type person struct {
	name   string
	age    int
	salary float64
}

func pointers() {
	fmt.Println("now we're doing pointers")
	i := 7
	inc(&i)
	fmt.Println(i)
}

func inc(x *int) {
	fmt.Println("original value of x: ", *x)
	*x++
}

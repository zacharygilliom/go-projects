package main

import (
	"fmt"
)

func add(a float64, b float64) float64 {
	return a + b
}

func multiply(c float64, d float64, e float64, f float64) float64 {
	num1 := add(c, d)
	num2 := add(e, f)
	product := num1 * num2
	return product

}

func main() {
	var num1 float64 = 5
	var num2 float64 = 10
	var num3 float64 = 2
	var num4 float64 = 11
	fmt.Println(multiply(num1, num2, num3, num4))
}

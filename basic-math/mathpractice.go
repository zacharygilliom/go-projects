package main

import (
	"fmt"
)

func add(a float64, b float64) float64 {
	return a + b
}

func divide(g float64, h float64) float64 {
	return g / h
}

func multiply(c float64, d float64) float64 {
	return c * d
}

func subtract(x float64, y float64) float64 {
	return x - y
}

func main() {
	var x1 float64 = 12.2
	var x2 float64 = 29.0
	k := add(x1, x2)
	j := subtract(x2, x1)
	m := multiply(k, j)
	n := divide(j, k)
	answer := multiply(m, n)
	fmt.Println(answer)
}

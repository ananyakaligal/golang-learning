package main

import "fmt"

func main() {
	// String
	var str string = "Golang"
	str2 := "Notes"

	// Integers
	var num int = 100
	var num8 int8 = 50
	var num64 int64 = 1000000

	// Floats
	var f32 float32 = 10.5
	var f64 float64 = 20.99

	// Boolean
	var b1 bool = true
	b2 := false

	// Print
	fmt.Println("Strings:", str, str2)
	fmt.Println("Integers:", num, num8, num64)
	fmt.Println("Floats:", f32, f64)
	fmt.Println("Booleans:", b1, b2)
}

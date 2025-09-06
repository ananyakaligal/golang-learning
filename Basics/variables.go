package main

import "fmt"

func main() {
	// Variables
	var name string = "xyz"
	name2 := "Ananya"

	var age int = 21
	age2 := 30

	var smallNum int8 = 100
	var largeNum int64 = 10000000000

	// Constants
	const PI = 3.14
	const Greeting = "Hello, Go!"

	// Booleans
	var isActive bool = true
	isActive2 := false

	// Print
	fmt.Println("Strings:", name, name2)
	fmt.Println("Integers:", age, age2, smallNum, largeNum)
	fmt.Println("Booleans:", isActive, isActive2)
	fmt.Println("Constants:", PI, Greeting)
}

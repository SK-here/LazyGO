package main

import (
	"fmt"
)

func main() {
	
	fmt.Println("Learn Go Programming")

	/*
	Go is a statically-typed language, meaning the type of every variable is known at compile-time. Common variable types include:
	Basic types: int, float64, bool, string
	*/

	// int
	var age int
	age = 24 // Var Value
	fmt.Println(age)

	// String
	var name string = "SK"
	fmt.Println(name)

	// Boolean Value
	FindOut := true
	fmt.Println(FindOut)

	// Float
	var pi float64 = 3.14159
	fmt.Println(pi)

	/*
	2. Multiple Variable Declarations;
	*/
	var a, b, c int
	var x, y = 1, 2
	a = 4 
	b = 6
	c = 8

	fmt.Println(x,y,a,b,c)

	/*
	Using short variable declaration 
	*/

	A := 42
	Name := "John"

	fmt.Println( A , Name)
}

package main

import (
	"fmt"
)

/*
A program which takes two inputs from the user and prints the sum of the two numbers.
*/

func add(a int, b int) int {
	return a + b
}

func main() {
	var a, b int

	fmt.Println("Enter a number: ")
	fmt.Scanln(&a)
	fmt.Println("Enter another number: ")
	fmt.Scanln(&b)

	fmt.Println("Sum of the two numbers is: ", add(a, b))
}

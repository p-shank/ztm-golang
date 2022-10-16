//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	greetPerson("shank")
	fmt.Println(returnMessage())

	var num1 int = returnNumber()
	var num2, num3 int = returnNumbers()

	fmt.Println("Number 1", num1)
	fmt.Println("Number 2", num2)
	fmt.Println("Number 3", num3)

	fmt.Println("Addition of numbers", add(num1, num2, num3))

}

func greetPerson(name string) {
	fmt.Println("Hi ", name)
}

func returnMessage() string {
	return "Hello"
}

func returnNumber() int {
	return rand.Intn(100)
}

func returnNumbers() (int, int) {
	return rand.Intn(100), rand.Intn(100)
}

func add(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

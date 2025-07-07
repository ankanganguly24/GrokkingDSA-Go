package main

import "fmt"

func main() {
	fmt.Println("Hello" + "World")
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(1+2, "is the sum of 1 and 2")
	// fmt.Println(1 + "2") // This line will throw an error , why? because we are trying to add a string to an integer which is not allowed in Go.
}

//GO has different values like strings, booleans, and integers. In this program, we are printing the concatenation of two strings, the result of a logical AND operation, the result of a logical OR operation, and the negation of a boolean value.

// The + operator is used to concatenate two strings. The && operator is used to perform a logical AND operation.

// The || operator is used to perform a logical OR operation. The ! operator is used to negate a boolean value. The output of this program will be: HelloWorld false true false 3 is the sum of 1 and 2

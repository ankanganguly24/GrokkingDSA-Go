// In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls

package main

import "fmt"

var q = "Heyaaa" // This is a package level variable. It can be accessed from any file in the same package. Here we see that q is a variable of type int and is assigned the value 1, which is an integer.

func main() {
	var a = "initial" // var declares 1 or more variables. Here we see that a is a variable of type string and is assigned the value "initial", which is a string.

	var b = 5 // You can declare multiple variables at once. Here we see that b is a variable of type int and is assigned the value 5, which is an integer.

	var c string // Go will infer the type of initialized variables. Here we see that c is a variable of type string and is assigned the value "cat", which is a string.

	c = "cat" // Go will infer the type of initialized variables. Here we see that c is a variable of type string and is assigned the value "cat", which is a string.

	var d int = 1 // Variables declared without a corresponding initialization are zero-valued. Here we see that d is a variable of type int and is assigned the value 1, which is an integer.

	var e bool // The zero value of a string is an empty string, the zero value of an int is 0, and the zero value of a bool is false.

	fmt.Println(a, b, c, d, e) // This will print the value of a, b, c, d, and e to the console.

	f := "apple" // The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.

	fmt.Println(f, q) // This will print the value of f to the console.

}

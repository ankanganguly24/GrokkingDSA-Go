// Go supports constants of character, string, boolean, and numeric values.

package main

import "fmt"

const s string = "constant" // This is how you declare a constant in Go. Here we see that s is a constant of type string and is assigned the value "constant", which is a string.

func main() {
	fmt.Printf(s) // This will print the value of s to the console.

	const d string = "constant2" // Constants can be declared like variables, but with the const keyword. Here we see that d is a constant of type string and is assigned the value "constant", which is a string.

	fmt.Println(d) // This will print the value of d to the console.
}

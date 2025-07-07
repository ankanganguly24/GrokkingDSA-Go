package main // package declaration basically tells Go that this file is part of the main package

import "fmt" // import statement is used to include the fmt package in the program which is used to print output to the console. fmt stands for format. There are many other packages in Go which provide different functionalities.

func main() { // main function is the entry point of any Go program. It is the first function that gets executed when you run the program.
	fmt.Println("Hello, World!")
	// fmt.Println is used to print output to the console. It adds a newline character at the end of the output.
} // closing brace of the main function

//To run the program, you can use the go run command followed by the name of the file. For example, to run the hello.go program, you can use the following command: go run hello.go

package main

import "fmt"

// Define a function type for the callback
type Callback func(string)

// A function that accepts a callback
func process(name string, callback Callback) {
	// Perform some processing
	result := "Hello, " + name
	// Invoke the callback with the result
	callback(result)
}

func main() {
	// Define a callback function
	myCallback := func(message string) {
		fmt.Println(message)
	}

	// Call the process function with a name and the callback
	process("World", myCallback)
}

package main

import "fmt"

func main() {
	var pointer1 *int
	var value1 = 12

	pointer1 = &value1
	doublePointerByValue(pointer1)

	fmt.Printf("The value of the original value is: %d\n", value1)
}

// Changing the value through a pointer:
// Write a function that takes a pointer to an int and
// doubles the value of that variable.
// Check that the original value is changed.
func doublePointerByValue(value *int) {
	valueOfPointer := *value

	*value = valueOfPointer * 2
}

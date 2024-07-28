package main

import "fmt"

// Stack defines a type for a stack using a slice of integers
type Stack []int

// Push adds an element to the top of the Stack
func (s *Stack) Push(value int) {
	*s = append(*s, value) // Append the given value to the end of the slice, which represents the top of the stack
}

// Pop removes and returns the top element of the Stack (LIFO: Last In, First Out)
func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return -1 // Return -1 if the Stack is empty, indicating there's nothing to pop
	}
	element := (*s)[len(*s)-1] // Retrieve the last element of the Stack (the top of the stack)
	*s = (*s)[:len(*s)-1]      // Remove the last element from the Stack by slicing the slice to exclude the last element
	return element             // Return the element that was removed
}

func main() {
	fmt.Println("STACK TESTS:")
	s := Stack{}                               // Initialize an empty Stack
	s.Push(1)                                  // Add the element 1 to the Stack
	s.Push(2)                                  // Add the element 2 to the Stack
	s.Push(3)                                  // Add the element 3 to the Stack
	fmt.Println("Stack: ", s)                  // Print the current state of the Stack
	fmt.Println("Stack Length: ", len(s))      // Print the number of elements in the Stack
	s.Pop()                                    // Remove the top element from the Stack
	fmt.Println("Stack: ", s)                  // Print the state of the Stack after popping
	fmt.Printf("Stack Length: %v\n\n", len(s)) // Print the number of elements in the Stack after popping
}

package main

// Stack starts as a simple slice
type Stack []int

// Push appends an element on the Stack
func (s *Stack) Push(value int) {
	*s = append(*s, value) // Appends the passed in parameter value on the Stack
}

// Pop removes the last element of the Stack (LIFO)
func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return -1 // if the Stack is empty exit the function
	}
	element := (*s)[len(*s)-1] // assign element to the last element in the Stack
	*s = (*s)[:len(*s)-1]      // Remove the last element in the Stack following (LIFO)
	return element             // return the element that has been removed from the Stack
}

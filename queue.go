package main

import "fmt"

// Queue defines a type for a queue using a slice of integers
type Queue []int

// Enqueue adds an element to the end of the Queue
func (q *Queue) Enqueue(value int) {
	*q = append(*q, value) // Append the given value to the end of the Queue
}

// Dequeue removes and returns the first element of the Queue (FIFO: First In, First Out)
func (q *Queue) Dequeue() int {
	if len(*q) == 0 {
		return -1 // Return -1 if the Queue is empty, indicating there's nothing to dequeue
	}
	element := (*q)[0] // Retrieve the first element of the Queue
	*q = (*q)[1:]      // Remove the first element from the Queue, slicing the slice to exclude the first element
	return element     // Return the element that was removed
}

func main() {
	fmt.Println("Queue Tests:")
	q := Queue{}                               // Initialize an empty Queue
	q.Enqueue(1)                               // Add the element 1 to the Queue
	q.Enqueue(2)                               // Add the element 2 to the Queue
	q.Enqueue(3)                               // Add the element 3 to the Queue
	fmt.Println("Queue: ", q)                  // Print the current state of the Queue
	fmt.Println("Queue Length: ", len(q))      // Print the number of elements in the Queue
	q.Dequeue()                                // Remove the first element from the Queue
	fmt.Println("Queue: ", q)                  // Print the state of the Queue after dequeueing
	fmt.Printf("Queue Length: %v\n\n", len(q)) // Print the number of elements in the Queue after dequeueing
}

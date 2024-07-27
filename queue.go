package main

// Queue starts as a simple slice
type Queue []int

// Enqueue appends an element on the Queue
func (q *Queue) Enqueue(value int) {
	*q = append(*q, value) // Appends the passed in parameter value on the Queue
}

// Dequeue removes the last element of the Queue (FIFO)
func (q *Queue) Dequeue() int {
	if len(*q) == 0 {
		return -1 // if the Queue is empty exit the function
	}
	element := (*q)[0] // assign element with the first element of the queue
	*q = (*q)[1:]      // remove the first element of the queue following (FIFO)
	return element     // return the element that was removed
}

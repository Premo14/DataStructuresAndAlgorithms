package main

import (
	"fmt"
)

// LinkedListNode defines a node within the linked list
type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

// LinkedList defines the linked list itself
type LinkedList struct {
	head *LinkedListNode
}

// Insert inserts a node to the end of the linked list
func (l *LinkedList) Insert(value int) {
	newNode := &LinkedListNode{value: value} // new node that holds the value passed to it
	if l.head == nil {
		l.head = newNode // if the head is empty make the new node the head
	} else {
		current := l.head // create current node for iterating over linked list
		for current.next != nil {
			current = current.next // iterate over linked list until the next node is empty
		}
		current.next = newNode // assign the new node to the end of the linked list
	}
}

// Delete deletes a node by the value passed by parameter
func (l *LinkedList) Delete(value int) {
	if l.Length() == 0 {
		return // if the linked list is empty, exit the function
	}

	toDelete := l.head // node to be deleted

	if l.head.value == value {
		l.head = l.head.next // if the head node value equals the value of the node to be deleted, assign the next node as the head
	}

	for toDelete.next.value != value {
		toDelete = toDelete.next // iterate over the linked list until the value to be deleted is found
	}
	toDelete.next = toDelete.next.next // delete the node with value to be deleted
}

// Display displays the linked list to the console
func (l *LinkedList) Display() {
	current := l.head
	fmt.Print("LinkedList: ")
	for current != nil {
		fmt.Println(current.value) // iterate over the linked list, printing each node until the next node is nil
		current = current.next
	}
}

// Length gets the length of the linked list
func (l *LinkedList) Length() int {
	length := 0       // counter
	current := l.head // assign current to the head of the linked list
	for current != nil {
		length++ // increment length until the next node is nil
		current = current.next
	}
	return length // return the length of the linked list
}

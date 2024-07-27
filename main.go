package main

import "fmt"

// Data Structures and Algorithms tests
func main() {
	// Data Structures ___________________________________________________________________
	// Linked List Tests -----------------------------------------------------------------
	fmt.Println("LINKED LIST TESTS:")
	ll := LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Display()
	fmt.Println("Length: ", ll.Length())
	ll.Delete(2)
	ll.Display()
	fmt.Printf("Linked List Length: %v\n\n", ll.Length())
	// End Linked List Tests ---------------------------------------------------------

	// Stack Tests ----------------------------------------------------------------------
	fmt.Println("STACK TESTS:")
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println("Stack: ", s)
	fmt.Println("Stack Length: ", len(s))
	s.Pop()
	fmt.Println("Stack: ", s)
	fmt.Printf("Stack Length: %v\n\n", len(s))
	// End Stack Tests ------------------------------------------------------------------

	// Queue Tests ----------------------------------------------------------------------
	fmt.Println("Queue Tests:")
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println("Queue: ", q)
	fmt.Println("Queue Length: ", len(q))
	q.Dequeue()
	fmt.Println("Queue: ", q)
	fmt.Printf("Queue Length: %v\n\n", len(q))
	// End Queue Tests ------------------------------------------------------------------

	// Binary Tree Tests ----------------------------------------------------------------
	root := &TreeNode{10, nil, nil}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(20)
	root.Display()
	fmt.Println("Binary Tree Length: ", root.Length())
	root.Delete(15)
	root.Delete(10)
	root.Display()
	fmt.Printf("Binary Tree Length: %v\n\n", root.Length())
	// End Binary Tree Tests ------------------------------------------------------------

	// Hash Table Tests -----------------------------------------------------------------
	h := HashTable{}
	h["one"] = 1
	h["two"] = 2
	h["three"] = 3
	fmt.Println("Hash Table: ", h)
	fmt.Printf("Hash Table Length: %v\n", len(h))
	delete(h, "two")
	fmt.Println("Hash Table: ", h)
	fmt.Printf("Hash Table Length: %v\n\n", len(h))
	// End Hash Table Tests -------------------------------------------------------------
	// End Data Structures ______________________________________________________________

	// Sorting Algorithms _______________________________________________________________
	// Bubble Sort ----------------------------------------------------------------------
	fmt.Println("Bubble Sort:")
	bubbleArr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original Array: ", bubbleArr)
	BubbleSort(bubbleArr)
	fmt.Printf("Bubble Sorted array: %v\n\n", bubbleArr)

	// Selection Sort -------------------------------------------------------------------
	fmt.Println("Selection Sort:")
	selectionArr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original Array: ", selectionArr)
	SelectionSort(selectionArr)
	fmt.Printf("Selection Sorted array: %v\n\n", selectionArr)

	// Insertion Sort -------------------------------------------------------------------
	fmt.Println("Insertion Sort:")
	insertionArr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original Array: ", insertionArr)
	InsertionSort(insertionArr)
	fmt.Printf("Insertion Sorted array: %v\n\n", insertionArr)

	// Merge Sort ----------------------------------------------------------------------
	fmt.Println("Merge Sort:")
	mergeArr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original Array: ", mergeArr)
	MergeSort(mergeArr, 0, len(mergeArr)-1)
	fmt.Printf("Merge Sorted array: %v\n\n", mergeArr)

	// Quick Sort ----------------------------------------------------------------------
	fmt.Println("Quick Sort:")
	quickArr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original Array: ", quickArr)
	QuickSort(quickArr, 0, len(quickArr)-1)
	fmt.Printf("Quick Sorted array: %v\n\n", quickArr)

	// Search Algorithms _______________________________________________________________
	sortedArr := []int{11, 12, 22, 25, 34, 64, 90}
	target := 25
	// Binary Search -------------------------------------------------------------------
	fmt.Println("Binary Search:")
	fmt.Println("Array: ", sortedArr)
	result := BinarySearch(sortedArr, target)
	found(target, result)

	// Linear Search -------------------------------------------------------------------
	fmt.Println("Linear Search:")
	fmt.Println("Array: ", sortedArr)
	result = LinearSearch(sortedArr, target)
	found(target, result)

	// Jump Search ---------------------------------------------------------------------
	fmt.Println("Jump Search:")
	fmt.Println("Array: ", sortedArr)
	result = JumpSearch(sortedArr, target)
	found(target, result)

	// Interpolation Search ------------------------------------------------------------
	fmt.Println("Interpolation Search:")
	fmt.Println("Array: ", sortedArr)
	result = InterpolationSearch(sortedArr, target)
	found(target, result)
	// End Algorithms __________________________________________________________________
}

// Function for generic algorithm tests above
func found(t, r int) {
	if r != -1 {
		fmt.Printf("%v found at index %d\n\n", t, r)
	} else {
		fmt.Printf("%v not found\n\n", t)
	}
}

package main

import "fmt"

// BubbleSort sorts an array by repeatedly stepping through the list,
// comparing adjacent elements, and swapping them if they are in the wrong order.
// The process is repeated until the array is sorted.
// Time Complexity: Best O(n), Avg O(n^2), Worst O(n^2)
func BubbleSort(arr []int) {
	n := len(arr)              // Get the number of elements in the array
	for i := 0; i < n-1; i++ { // Iterate through the array
		for j := 0; j < n-i-1; j++ { // Compare adjacent elements up to the sorted section
			if arr[j] > arr[j+1] { // If the current element is greater than the next element
				// Swap the elements to place the larger one at the end
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// SelectionSort sorts an array by repeatedly finding the minimum element from the unsorted part
// and putting it at the beginning of the array.
// Time Complexity: Best O(n^2), Avg O(n^2), Worst O(n^2)
func SelectionSort(arr []int) {
	n := len(arr)              // Get the number of elements in the array
	for i := 0; i < n-1; i++ { // Iterate through the array
		minIndex := i                // Assume the current index is the minimum
		for j := i + 1; j < n; j++ { // Find the minimum element in the unsorted part
			if arr[j] < arr[minIndex] { // If the current element is smaller than the assumed minimum
				minIndex = j // Update the minimum index
			}
		}
		// Swap the found minimum element with the first element of the unsorted part
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// InsertionSort builds the final sorted array one item at a time by
// comparing each item with the already-sorted portion and inserting it into the correct position.
// Time Complexity: Best O(n), Avg O(n^2), Worst O(n^2)
func InsertionSort(arr []int) {
	n := len(arr)            // Get the number of elements in the array
	for i := 1; i < n; i++ { // Iterate through the array starting from the second element
		key := arr[i] // Current element to be inserted into the sorted portion
		j := i - 1
		// Move elements of the sorted portion that are greater than the key to one position ahead
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		// Insert the key at its correct position
		arr[j+1] = key
	}
}

// Merge merges two sorted sub-arrays into a single sorted array.
// Part of the MergeSort algorithm.
func Merge(arr []int, l, m, r int) {
	// Calculate the sizes of the two sub-arrays to be merged
	n1 := m - l + 1
	n2 := r - m

	// Create temporary arrays to hold the values of the two sub-arrays
	L := make([]int, n1)
	R := make([]int, n2)

	// Copy data to temporary arrays L[] and R[]
	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	// Merge the temporary arrays back into the original array
	i := 0 // Initial index of the first sub-array
	j := 0 // Initial index of the second sub-array
	k := l // Initial index of the merged array
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i] // Place the smaller element into the original array
			i++
		} else {
			arr[k] = R[j] // Place the smaller element into the original array
			j++
		}
		k++
	}

	// Copy the remaining elements of L[], if any
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	// Copy the remaining elements of R[], if any
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

// MergeSort is a recursive function that sorts an array using the Merge Sort algorithm.
// It divides the array into halves, recursively sorts each half, and then merges the sorted halves.
// Time Complexity: Best O(nlog_n), Avg O(nlog_n), Worst O(nlog_n)
func MergeSort(arr []int, l, r int) {
	if l < r { // Base case: if the sub-array has more than one element
		m := l + (r-l)/2 // Find the middle point to divide the array into two halves

		// Recursively sort the two halves
		MergeSort(arr, l, m)
		MergeSort(arr, m+1, r)

		// Merge the sorted halves
		Merge(arr, l, m, r)
	}
}

// Partition partitions the array around a pivot element
// and is used as part of the QuickSort algorithm.
// The pivot element is placed in its correct position, and the elements are rearranged
// such that elements smaller than the pivot come before it and elements greater come after it.
func Partition(arr []int, low, high int) int {
	pivot := arr[high] // Choose the pivot element (last element in this case)
	i := low - 1       // Index of the smaller element

	for j := low; j < high; j++ { // Iterate through the array
		if arr[j] < pivot { // If the current element is smaller than the pivot
			i++                             // Increment the index of the smaller element
			arr[i], arr[j] = arr[j], arr[i] // Swap the elements
		}
	}
	// Swap the pivot element to its correct position
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1 // Return the index of the pivot element
}

// QuickSort selects a pivot element, partitions the array around the pivot,
// and recursively sorts the sub-arrays on either side of the pivot.
// Time Complexity: Best O(nlog_n), Avg O(nlog_n), Worst O(n^2)
func QuickSort(arr []int, low, high int) {
	if low < high { // Base case: if the sub-array has more than one element
		pi := Partition(arr, low, high) // Partition the array and get the pivot index
		QuickSort(arr, low, pi-1)       // Recursively sort the left part
		QuickSort(arr, pi+1, high)      // Recursively sort the right part
	}
}

func main() {
	// Bubble Sort ----------------------------------------------------------------------
	fmt.Println("Bubble Sort:")
	bubbleArr := []int{64, 34, 25, 12, 22, 11, 90}       // Initialize an array
	fmt.Println("Original Array: ", bubbleArr)           // Print the original array
	BubbleSort(bubbleArr)                                // Sort the array using Bubble Sort
	fmt.Printf("Bubble Sorted array: %v\n\n", bubbleArr) // Print the sorted array

	// Selection Sort -------------------------------------------------------------------
	fmt.Println("Selection Sort:")
	selectionArr := []int{64, 34, 25, 12, 22, 11, 90}          // Initialize an array
	fmt.Println("Original Array: ", selectionArr)              // Print the original array
	SelectionSort(selectionArr)                                // Sort the array using Selection Sort
	fmt.Printf("Selection Sorted array: %v\n\n", selectionArr) // Print the sorted array

	// Insertion Sort -------------------------------------------------------------------
	fmt.Println("Insertion Sort:")
	insertionArr := []int{64, 34, 25, 12, 22, 11, 90}          // Initialize an array
	fmt.Println("Original Array: ", insertionArr)              // Print the original array
	InsertionSort(insertionArr)                                // Sort the array using Insertion Sort
	fmt.Printf("Insertion Sorted array: %v\n\n", insertionArr) // Print the sorted array

	// Merge Sort ----------------------------------------------------------------------
	fmt.Println("Merge Sort:")
	mergeArr := []int{64, 34, 25, 12, 22, 11, 90}      // Initialize an array
	fmt.Println("Original Array: ", mergeArr)          // Print the original array
	MergeSort(mergeArr, 0, len(mergeArr)-1)            // Sort the array using Merge Sort
	fmt.Printf("Merge Sorted array: %v\n\n", mergeArr) // Print the sorted array

	// Quick Sort ----------------------------------------------------------------------
	fmt.Println("Quick Sort:")
	quickArr := []int{64, 34, 25, 12, 22, 11, 90}      // Initialize an array
	fmt.Println("Original Array: ", quickArr)          // Print the original array
	QuickSort(quickArr, 0, len(quickArr)-1)            // Sort the array using Quick Sort
	fmt.Printf("Quick Sorted array: %v\n\n", quickArr) // Print the sorted array
}

package main

// BubbleSort repeatedly steps through the list, compares adjacent elements, and swaps
// them if they are in the wrong order.
// Best Case O(n), Avg Case O(n^2), Worst Case O(n^2)
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ { // Loop over each element in the array
		for j := 0; j < n-i-1; j++ { // Compare each pair of adjacent elements
			if arr[j] > arr[j+1] { // If elements are in wrong order
				// Swap the elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// SelectionSort repeatedly finds the minimum element from the unsorted
// part and puts it at the beginning.
// Best Case O(n^2), Avg Case O(n^2), Worst Case O(n^2)
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i // Assume the current index is the minimum
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] { // Find the minimum element
				minIndex = j
			}
		}
		// Swap the found minimum element with the first element
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// InsertionSort builds the final sorted array one item at a time by
// comparing each item with the already-sorted portion.
// Best Case O(n), Avg Case O(n^2), Worst Case O(n^2)
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i] // Current element to be inserted
		j := i - 1
		// Move elements of arr[0...i-1] that are greater than key to one position ahead
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key // Insert the key at the correct position
	}
}

// Merge merges two sub-arrays of arr[]
// Part of MergeSort
func Merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	// Create temporary arrays
	L := make([]int, n1)
	R := make([]int, n2)

	// Copy data to temporary arrays L[] and R[]
	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	// Merge the temporary arrays back into arr[l...r]
	i := 0
	j := 0
	k := l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
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

// MergeSort is a recursive function that sorts the array using Merge Sort algorithm
// Best Case O(nlog_n), Avg Case O(nlog_n), Worst Case O(log_n)
func MergeSort(arr []int, l, r int) {
	if l < r {
		m := l + (r-l)/2

		// Sort first and second halves
		MergeSort(arr, l, m)
		MergeSort(arr, m+1, r)

		// Merge the sorted halves
		Merge(arr, l, m, r)
	}
}

// Partition partitions the array around the pivot
// Part of QuickSort
func Partition(arr []int, low, high int) int {
	pivot := arr[high] // Pivot element
	i := low - 1       // Index of smaller element

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // Swap if element is smaller than pivot
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Swap pivot element to the correct position
	return i + 1
}

// QuickSort selects a 'pivot' element, partitions the array around the pivot, and
// recursively sorts the sub-arrays.
// Best Case O(nlog_n), Avg Case O(nlog_n), Worst Case O(n^2)
func QuickSort(arr []int, low, high int) {
	if low < high {
		pi := Partition(arr, low, high) // Partition the array
		QuickSort(arr, low, pi-1)       // Recursively sort the left part
		QuickSort(arr, pi+1, high)      // Recursively sort the right part
	}
}

package main

import (
	"fmt"
	"math"
)

// BinarySearch efficiently finds an item in a sorted array by repeatedly dividing
// the search interval in half.
// Time Complexity: Best O(1), Avg O(log_n), Worst O(log_n)
func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1 // Initialize the search interval
	for low <= high {          // Continue searching while the interval is valid
		mid := low + (high-low)/2 // Calculate the middle index
		if arr[mid] == target {   // If the target is found
			return mid // Return the index
		}
		if arr[mid] < target { // If the target is in the right half
			low = mid + 1 // Narrow the interval to the right half
		} else { // If the target is in the left half
			high = mid - 1 // Narrow the interval to the left half
		}
	}
	return -1 // Return -1 if the target is not found
}

// LinearSearch sequentially checks each element of the array until the target
// value is found or the end of the array is reached.
// Time Complexity: Best O(1), Avg O(n), Worst O(n)
func LinearSearch(arr []int, target int) int {
	for i, value := range arr { // Loop through each element in the array
		if value == target { // If the current element is the target
			return i // Return the index of the target
		}
	}
	return -1 // Return -1 if the target is not found
}

// JumpSearch searches for a value by jumping ahead by fixed steps and then
// performing a linear search within a block.
// Time Complexity: Best O(1), Avg O(sqrt(n)), Worst O(sqrt(n))
func JumpSearch(arr []int, target int) int {
	n := len(arr)                      // Get the number of elements in the array
	step := int(math.Sqrt(float64(n))) // Calculate the jump size based on the square root of the array size
	prev := 0                          // Initial position for the previous block

	// Jump ahead by fixed steps until the end of the array or until a block with the target
	for arr[int(math.Min(float64(step), float64(n)))-1] < target {
		prev = step                        // Move to the next block
		step += int(math.Sqrt(float64(n))) // Move the step forward
		if prev >= n {                     // If the previous block goes out of bounds
			return -1 // Return -1 if the target is not found
		}
	}

	// Perform a linear search within the block where the target could be
	for arr[prev] < target {
		prev++                                                // Move to the next element in the block
		if prev == int(math.Min(float64(step), float64(n))) { // If the end of the block is reached
			return -1 // Return -1 if the target is not found
		}
	}

	if arr[prev] == target { // If the target is found
		return prev // Return the index of the target
	}

	return -1 // Return -1 if the target is not found
}

// InterpolationSearch is an improvement over binary search for uniformly distributed
// arrays. It estimates the position of the target value.
// Time Complexity: Best O(1), Avg O(loglog_n), Worst O(n)
func InterpolationSearch(arr []int, target int) int {
	low := 0             // Start of the search interval
	high := len(arr) - 1 // End of the search interval

	// Continue searching while the interval is valid and the target is within the range
	for low <= high && target >= arr[low] && target <= arr[high] {
		if low == high { // If there is only one element left
			if arr[low] == target {
				return low // Return the index if the target is found
			}
			return -1 // Return -1 if the target is not found
		}

		// Estimate the position of the target based on the distribution
		pos := low + int(float64(high-low)*(float64(target-arr[low])/float64(arr[high]-arr[low])))

		if arr[pos] == target { // If the target is found at the estimated position
			return pos // Return the index of the target
		}

		if arr[pos] < target { // If the target is in the right sub-array
			low = pos + 1 // Narrow the interval to the right half
		} else { // If the target is in the left sub-array
			high = pos - 1 // Narrow the interval to the left half
		}
	}

	return -1 // Return -1 if the target is not found
}

func main() {
	sortedArr := []int{11, 12, 22, 25, 34, 64, 90} // Example sorted array
	target := 25                                   // Target value to search for

	// Binary Search -------------------------------------------------------------------
	fmt.Println("Binary Search:")
	fmt.Println("Array: ", sortedArr)
	result := BinarySearch(sortedArr, target) // Perform binary search
	found(target, result)                     // Print the result

	// Linear Search -------------------------------------------------------------------
	fmt.Println("Linear Search:")
	fmt.Println("Array: ", sortedArr)
	result = LinearSearch(sortedArr, target) // Perform linear search
	found(target, result)                    // Print the result

	// Jump Search ---------------------------------------------------------------------
	fmt.Println("Jump Search:")
	fmt.Println("Array: ", sortedArr)
	result = JumpSearch(sortedArr, target) // Perform jump search
	found(target, result)                  // Print the result

	// Interpolation Search ------------------------------------------------------------
	fmt.Println("Interpolation Search:")
	fmt.Println("Array: ", sortedArr)
	result = InterpolationSearch(sortedArr, target) // Perform interpolation search
	found(target, result)                           // Print the result
}

// Function for generic algorithm tests above
func found(t, r int) {
	if r != -1 {
		fmt.Printf("%v found at index %d\n\n", t, r) // Print the index if the target is found
	} else {
		fmt.Printf("%v not found\n\n", t) // Print a not found message if the target is not found
	}
}

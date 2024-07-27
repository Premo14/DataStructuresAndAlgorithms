package main

import "math"

// BinarySearch efficiently finds an item in a sorted array by repeatedly dividing
// the search interval in half.
// Best Case O(1), Avg Case O(log_n), Worst Case O(log_n)
func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2 // Find the middle index
		if arr[mid] == target {
			return mid // Return the index if the target is found
		}
		if arr[mid] < target {
			low = mid + 1 // Search in the right half
		} else {
			high = mid - 1 // Search in the left half
		}
	}
	return -1 // Return -1 if the target is not found
}

// LinearSearch sequentially checks each element of the array until the target
// value is found or the end of the array is reached.
// Best Case O(1), Avg Case O(n), Worst Case O(n)
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
// Best Case O(1), Avg Case O(sqrt(n)), Worst Case O(sqrt(n))
func JumpSearch(arr []int, target int) int {
	n := len(arr)
	step := int(math.Sqrt(float64(n))) // Calculate the jump size
	prev := 0

	for arr[int(math.Min(float64(step), float64(n)))-1] < target {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1 // Return -1 if the target is not found
		}
	}

	for arr[prev] < target {
		prev++
		if prev == int(math.Min(float64(step), float64(n))) {
			return -1 // Return -1 if the target is not found
		}
	}

	if arr[prev] == target {
		return prev // Return the index of the target
	}

	return -1 // Return -1 if the target is not found
}

// InterpolationSearch is an improvement over binary search for uniformly distributed
// arrays. It estimates the position of the target value.
// Best Case O(1), Avg Case O(loglog_n), Worst Case O(n)
func InterpolationSearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high && target >= arr[low] && target <= arr[high] {
		if low == high {
			if arr[low] == target {
				return low
			}
			return -1
		}

		// Estimate the position
		pos := low + int(float64(high-low)*(float64(target-arr[low])/float64(arr[high]-arr[low])))

		if arr[pos] == target {
			return pos // Return the index of the target
		}

		if arr[pos] < target {
			low = pos + 1 // Target is in the right subarray
		} else {
			high = pos - 1 // Target is in the left subarray
		}
	}

	return -1 // Return -1 if the target is not found
}

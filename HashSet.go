package main

import "fmt"

// HashSet represents a set of unique elements.
// Internally, it uses a map where keys are elements and values are empty structs.
type HashSet map[string]struct{}

// NewHashSet creates and returns a new instance of HashSet.
// It initializes the underlying map to store elements.
func NewHashSet() HashSet {
	// Initialize and return a new HashSet with an empty map
	return make(HashSet)
}

// Add inserts a new element into the HashSet.
// If the element already exists, it remains unchanged.
func (hs HashSet) Add(element string) {
	// Add the element to the map with an empty struct value.
	// The empty struct is used because it requires zero memory.
	hs[element] = struct{}{}
}

// Remove deletes an element from the HashSet.
// If the element does not exist, the operation has no effect.
func (hs HashSet) Remove(element string) {
	// Delete the element from the map.
	// If the element is not present, this operation is harmless.
	delete(hs, element)
}

// Contains checks if an element is present in the HashSet.
// It returns true if the element exists, otherwise false.
func (hs HashSet) Contains(element string) bool {
	// Check if the element exists in the map.
	// The second value in the map lookup indicates if the key was found.
	_, exists := hs[element]
	return exists // Return true if the element exists, false otherwise
}

func main() {
	// Create a new HashSet instance
	hs := NewHashSet()

	// Add elements to the HashSet
	hs.Add("apple")
	hs.Add("banana")

	// Check if elements are present in the HashSet
	fmt.Println(hs.Contains("apple"))  // Output: true
	fmt.Println(hs.Contains("orange")) // Output: false

	// Remove an element from the HashSet
	hs.Remove("apple")

	// Check if the element was successfully removed
	fmt.Println(hs.Contains("apple")) // Output: false
}

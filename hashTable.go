package main

import "fmt"

// Go has a HashTable data structure built in called a map

type HashTable = map[string]int

func main() {
	h := HashTable{}                                // Create an empty HashTable object (really just a alias for a map)
	h["one"] = 1                                    // Add the key "one" and pair it to int 1
	h["two"] = 2                                    // Add the key "two" and pair it to int 2
	h["three"] = 3                                  // Add the key "three" and pair it to int 3
	fmt.Println("Hash Table: ", h)                  // Display the contents of the map
	fmt.Printf("Hash Table Length: %v\n", len(h))   // Display the length of the map
	delete(h, "two")                                // Delete an element out of the map with key "two"
	fmt.Println("Hash Table: ", h)                  // Display the updated contents of the map
	fmt.Printf("Hash Table Length: %v\n\n", len(h)) // Display the updated length of the map
}

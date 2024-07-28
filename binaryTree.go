package main

import "fmt"

// TreeNode defines a structure for a node in a binary tree
type TreeNode struct {
	value int       // The value stored in the node
	left  *TreeNode // Pointer to the left child node
	right *TreeNode // Pointer to the right child node
}

// Insert adds a new node with the given value to the binary tree
func (t *TreeNode) Insert(value int) {
	if t == nil {
		t.value = value // If the tree is empty (root is nil), set the node value
		return
	}
	if value <= t.value { // If the value to insert is less than or equal to the current node's value
		if t.left == nil { // If there is no left child
			t.left = &TreeNode{value: value} // Create a new node and set it as the left child
		} else {
			t.left.Insert(value) // Recursively insert the value into the left subtree
		}
	} else { // If the value to insert is greater than the current node's value
		if t.right == nil { // If there is no right child
			t.right = &TreeNode{value: value} // Create a new node and set it as the right child
		} else {
			t.right.Insert(value) // Recursively insert the value into the right subtree
		}
	}
}

// Delete removes a node with the given value from the binary tree
func (t *TreeNode) Delete(value int) *TreeNode {
	if t == nil {
		return nil // If the tree is empty, there is nothing to delete
	}

	if value < t.value { // If the value to delete is less than the current node's value
		t.left = t.left.Delete(value) // Recursively delete the value from the left subtree
	} else if value > t.value { // If the value to delete is greater than the current node's value
		t.right = t.right.Delete(value) // Recursively delete the value from the right subtree
	} else {
		// Node with the value to delete is found

		// Case 1: Node has no children
		if t.left == nil && t.right == nil {
			return nil // Remove the node by returning nil
		}

		// Case 2: Node has one child
		if t.left == nil {
			return t.right // Replace the node with its right child
		}
		if t.right == nil {
			return t.left // Replace the node with its left child
		}

		// Case 3: Node has two children
		minRight := t.right.findMin()            // Find the smallest value in the right subtree
		t.value = minRight.value                 // Replace the current node's value with the smallest value
		t.right = t.right.Delete(minRight.value) // Delete the smallest value from the right subtree
	}

	return t // Return the updated node
}

// findMin returns the node with the smallest value in the subtree
func (t *TreeNode) findMin() *TreeNode {
	current := t
	for current.left != nil { // Traverse to the leftmost node
		current = current.left
	}
	return current // Return the node with the minimum value
}

// Display prints the structure of the binary tree
func (t *TreeNode) Display() {
	printTree(t, "", true) // Call printTree to display the tree structure
}

// printTree is a helper function that prints the binary tree in a readable format
func printTree(t *TreeNode, prefix string, isTail bool) {
	if t != nil {
		fmt.Print(prefix) // Print the prefix (indentation)
		if isTail {       // If the node is a tail (i.e., it has no sibling)
			fmt.Print("└── ") // Print the tail symbol
			prefix += "    "  // Update prefix for children nodes
		} else {
			fmt.Print("├── ") // Print the branch symbol
			prefix += "│   "  // Update prefix for children nodes
		}
		fmt.Println(t.value)             // Print the value of the current node
		printTree(t.left, prefix, false) // Recursively print the left subtree
		printTree(t.right, prefix, true) // Recursively print the right subtree
	}
}

// Length calculates the total number of nodes in the binary tree
func (t *TreeNode) Length() int {
	if t == nil {
		return 0 // If the tree is empty, return 0
	}
	return 1 + t.left.Length() + t.right.Length() // Return 1 (for the current node) plus the length of left and right subtrees
}

func main() {
	root := &TreeNode{10, nil, nil}                         // Create the root node with value 10
	root.Insert(5)                                          // Insert value 5 into the tree
	root.Insert(15)                                         // Insert value 15 into the tree
	root.Insert(3)                                          // Insert value 3 into the tree
	root.Insert(20)                                         // Insert value 20 into the tree
	root.Display()                                          // Display the tree structure
	fmt.Println("Binary Tree Length: ", root.Length())      // Print the total number of nodes in the tree
	root.Delete(15)                                         // Delete the node with value 15 from the tree
	root.Delete(10)                                         // Delete the node with value 10 from the tree
	root.Display()                                          // Display the updated tree structure
	fmt.Printf("Binary Tree Length: %v\n\n", root.Length()) // Print the updated total number of nodes in the tree
}

package main

import "fmt"

// TreeNode creates a binary tree struct that acts as a parent and points to two children (left and right)
type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// Insert inserts a new node into the binary tree
func (t *TreeNode) Insert(value int) {
	if t == nil {
		t.value = value // if binary tree is empty assign the root value to the value to be added
	}
	if value <= t.value { // if the new node value is less than or equal to the current node value
		if t.left == nil { // if the left child node is nil
			t.left = &TreeNode{value: value} // assign the left child node the new node value
		} else {
			t.left.Insert(value) // if the left child node is not nil, make the left child node current and retry this function
		}
	} else { // if the new node value is greater than the current node value
		if t.right == nil { // if the right child node is empty
			t.right = &TreeNode{value: value} // assign the right child node the new node value
		} else {
			t.right.Insert(value) // if the right child node is not nil, make the right child node current and retry this function
		}
	}
}

// Delete deletes a node in the binary tree with a given value
func (t *TreeNode) Delete(value int) *TreeNode {
	if t == nil {
		return nil // if the binary tree is empty, exit the function
	}

	if value < t.value { // if the value to be deleted is less than the current node value
		t.left = t.left.Delete(value) // try the function again with the left child node value
	} else if value > t.value { // if the value to be deleted is greater than the current node
		t.right = t.right.Delete(value) // try the function again with the right child node
	} else {
		// Node with the value found

		// Case 1: Node has no children
		if t.left == nil && t.right == nil {
			return nil
		}

		// Case 2: Node has one child
		if t.left == nil {
			return t.right
		}
		if t.right == nil {
			return t.left
		}

		// Case 3: Node has two children
		minRight := t.right.findMin()
		t.value = minRight.value
		t.right = t.right.Delete(minRight.value)
	}

	return t
}

// findMin finds the node with the minimum value in the subtree
func (t *TreeNode) findMin() *TreeNode {
	current := t
	for current.left != nil {
		current = current.left
	}
	return current
}

// Display displays the binary tree structure
func (t *TreeNode) Display() {
	printTree(t, "", true) // Calls printTree() to accurately show tree structure
}

// printTree is a helper function to print the binary tree
func printTree(t *TreeNode, prefix string, isTail bool) {
	if t != nil {
		fmt.Print(prefix)
		if isTail {
			fmt.Print("└── ")
			prefix += "    "
		} else {
			fmt.Print("├── ")
			prefix += "│   "
		}
		fmt.Println(t.value)
		printTree(t.left, prefix, false)
		printTree(t.right, prefix, true)
	}
}

// Length computes the number of nodes in the binary tree
func (t *TreeNode) Length() int {
	if t == nil {
		return 0 // if the tree is empty, exit the function
	}
	return 1 + t.left.Length() + t.right.Length() // returns the number of nodes in the binary tree
}

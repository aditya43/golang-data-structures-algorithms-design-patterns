package main

import "fmt"

var interations int

type Tree struct {
	node *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

// Tree's Insert function
func (t *Tree) Insert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.Insert(value)
	}

	return t // Returning Tree so we can use chained functions
}

// Node's Insert function
func (n *Node) Insert(value int) {
	if value <= n.value {
		// Go to left
		if n.left == nil {
			// If there is no node at left side, create here..
			n.left = &Node{value: value}
		} else {
			// If there is a node at left, call insert..
			n.left.Insert(value) // Calling Insert func recursively
		}
	} else {
		// Go to right
		if n.right == nil {
			// If there is no node at right side, create here..
			n.right = &Node{value: value}
		} else {
			// If there is a node at right, call insert..
			n.right.Insert(value) // Calling Insert func recursively
		}
	}
}

// Print nodes recursively
func printNode(n *Node) {
	if n == nil {
		return
	}

	fmt.Printf("   %d   ", n.value)
	printNode(n.left)  // Recursively call printNode
	printNode(n.right) // Recursively call printNode
}

// Search element using Binary Search
func (n *Node) Search(value int) {
	if n == nil {
		fmt.Print("\n\n+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+\n")
		fmt.Printf("Value to find: %d\n", value)
		fmt.Printf("Result: %d doesn't exists in Binary Tree\n", value)
		fmt.Printf("Iterations Took: %d", interations)
		fmt.Print("\n+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+\n")
		return
	}

	interations++

	if n.value == value {
		fmt.Print("\n\n+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+\n")
		fmt.Printf("Value to find: %d\n", value)
		fmt.Printf("Result: %d exists in Binary Tree\n", value)
		fmt.Printf("Iterations Took: %d", interations)
		fmt.Print("\n+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+\n")
		return
	}

	if value <= n.value {
		// Search in left
		n.left.Search(value) // Recursively call Search function
	} else {
		// Search in right
		n.right.Search(value) // Recursively call Search function
	}
}

func main() {
	t := &Tree{}

	t.Insert(10).Insert(8).Insert(20).Insert(11).Insert(0).Insert(30)

	fmt.Print("\n\nFull Binary Tree: ")
	printNode(t.node) // Print Binrary Tree

	interations = 0
	t.node.Search(0) // Search for value

	interations = 0
	t.node.Search(10) // Search for value

	interations = 0
	t.node.Search(300) // Search for value
}

package main

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

	println(n.value)
	printNode(n.left)  // Recursively call printNode
	printNode(n.right) // Recursively call printNode
}

func main() {
	t := &Tree{}

	t.Insert(10).Insert(8).Insert(20).Insert(11).Insert(0).Insert(30)

	printNode(t.node)
}

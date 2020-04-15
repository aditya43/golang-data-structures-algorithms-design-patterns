// Doubly Linked List
package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	value string
	next  *Node
	prev  *Node
}

// Get first node
func (l *List) First() *Node {
	return l.head
}

func (l *List) Tail() *Node {
	return l.tail
}

// Add new node at the end of the list. i.e. tail
func (l *List) Push(value string) {
	node := &Node{value: value}

	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}

	l.tail = node
}

// Get next node
func (n *Node) Next() *Node {
	return n.next
}

// Get previous node
func (n *Node) Prev() *Node {
	return n.prev
}

func main() {
	l := &List{}

	l.Push("Adi")
	l.Push("Nishi")
	l.Push("John")
	l.Push("Jane")
	l.Push("Foo")

	n := l.First()
	t := l.Tail()

	fmt.Printf("\nFirst node value: %s\n", n.value)
	fmt.Printf("Tail node value: %s\n", t.value)
	fmt.Print("\n-------------FULL LINKED LIST - FORWARD-------------\n")
	for {
		fmt.Printf("%s", n.value)
		n = n.Next()

		if n == nil {
			break
		} else {
			fmt.Print(" --> ")
		}
	}

	fmt.Print("\n------------FULL LINKED LIST - BACKWARD-------------\n")

	for {
		fmt.Printf("%s", t.value)
		t = t.Prev()

		if t == nil {
			break
		} else {
			fmt.Print(" <-- ")
		}
	}

	fmt.Print("\n----------------------------------------------------\n")
}

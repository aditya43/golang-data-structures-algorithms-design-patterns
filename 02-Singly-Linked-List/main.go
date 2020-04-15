// Singly Linked List
package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	value string
	next  *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Push(value string) {
	node := &Node{value: value}

	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}

	l.tail = node
}

func (n *Node) Next() *Node {
	return n.next
}

func main() {
	l := &List{}

	l.Push("Adi")
	l.Push("Nishi")
	l.Push("John")
	l.Push("Jane")

	n := l.First()

	for {
		fmt.Printf("%s", n.value)
		n = n.Next()

		if n == nil {
			break
		} else {
			fmt.Print(" --> ")
		}
	}
}

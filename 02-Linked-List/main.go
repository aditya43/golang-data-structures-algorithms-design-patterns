package main

type List struct {
	head *Node
}

type Node struct {
	value string
	next  *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Push(value string) {

}

func main() {

}

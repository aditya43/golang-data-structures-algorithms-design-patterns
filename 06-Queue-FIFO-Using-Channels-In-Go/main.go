// Queue - FIFO - Using channels
package main

import "fmt"

type Queue struct {
	items chan string
}

func (q *Queue) Enqueue(val string) {
	q.items <- val
}

func (q *Queue) Dequeue() string {
	return <-q.items
}

func main() {
	q := Queue{
		items: make(chan string, 16),
	}

	q.Enqueue("Adi")
	q.Enqueue("Nishi")
	q.Enqueue("John")
	q.Enqueue("Jane")

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	// fmt.Println(q.Dequeue()) // Trying to Dequeue non existent item
}

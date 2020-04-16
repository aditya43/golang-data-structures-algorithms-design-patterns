// Queue - FIFO - Using Slices
package main

import "fmt"

type Queue struct {
	items []string
}

func (q *Queue) Enqueue(val string) {
	q.items = append(q.items, val)
}

func (q *Queue) Dequeue() string {
	if len(q.items) == 0 {
		return "0 items left Queue to dequeue"
	}

	item, items := q.items[0], q.items[1:]

	q.items = items

	return item
}

func main() {
	q := Queue{}

	q.Enqueue("Adi")
	q.Enqueue("Nishi")
	q.Enqueue("John")
	q.Enqueue("Jane")

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue()) // Trying to Dequeue non existent item
}

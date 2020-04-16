// Stack
package main

import "fmt"

type Stack struct {
	items []string
}

func (s *Stack) Push(value string) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() string {
	front := len(s.items)

	if front < 1 {
		return "0 items left in stack to pop"
	}

	item, items := s.items[front-1], s.items[0:front-1]
	s.items = items
	return item
}

func main() {
	s := Stack{}

	s.Push("Adi")
	s.Push("Nishi")
	s.Push("John")
	s.Push("Jane")

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop()) // Popping non existent item
}

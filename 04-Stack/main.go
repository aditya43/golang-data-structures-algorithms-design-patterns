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
	left := len(s.items)

	if left < 1 {
		return "No more elements in stack"
	}

	item, items := s.items[left-1], s.items[0:left-1]
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

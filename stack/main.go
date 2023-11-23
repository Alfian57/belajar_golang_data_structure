package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)

}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("Cannot pop the stack because stack is empty")
		return -1
	}

	length := len(s.items) - 1
	toRemove := s.items[length]
	s.items = s.items[:length]

	return toRemove
}

func main() {
	myStack := &Stack{}
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)

	fmt.Println(myStack)

	myStack.Pop()

	fmt.Println(myStack)
}

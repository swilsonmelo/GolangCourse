package main

import "fmt"

type Stack struct {
	items []int
}

//push
func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

//pop
func (s *Stack) pop() int {
	stackLen := len(s.items) - 1
	toRemove := s.items[stackLen]
	s.items = s.items[:stackLen]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.push(100)
	myStack.push(200)
	myStack.push(300)
	fmt.Println(myStack)
	myStack.pop()
	fmt.Println(myStack)

}

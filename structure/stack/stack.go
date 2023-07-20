package stack

import (
	"errors"
	"fmt"
)

var ErrStackIsEmpty = errors.New("stack is empty")

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)

	fmt.Printf("data: %v, len: %d, cap: %d \n\r", s.items, len(s.items), cap(s.items))
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrStackIsEmpty
	}

	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	fmt.Printf("data: %v, len: %d, cap: %d \n\r", s.items, len(s.items), cap(s.items))

	return val, nil
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, ErrStackIsEmpty
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Render() {
	fmt.Println(s.items)
}

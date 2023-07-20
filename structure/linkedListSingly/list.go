package linkedListSingly

import "errors"

var ErrSinglyLinkedListIsEmpty = errors.New("singly linked list is empty")

type SinglyLinkedList struct {
	Count int
	Head  *Node
	Tail  *Node
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (s *SinglyLinkedList) AddFirst(val int) {
	node := NewNode(val)

	if s.IsEmpty() {
		s.Head = node
		s.Tail = node
	} else {
		node.Next = s.Head
		s.Head = node
	}

	s.Count++
}

func (s *SinglyLinkedList) AddLast(val int) {
	node := NewNode(val)

	if s.IsEmpty() {
		s.Head = node
		s.Tail = node
	} else {
		s.Tail.Next = node
		s.Tail = node
	}

	s.Count++
}

func (s *SinglyLinkedList) RemoveFirst() error {
	if s.IsEmpty() {
		return ErrSinglyLinkedListIsEmpty
	}

	s.Head = s.Head.Next

	s.Count--

	if s.IsEmpty() {
		s.Tail = nil
	}

	return nil
}

func (s *SinglyLinkedList) RemoveLast() error {
	if s.IsEmpty() {
		return ErrSinglyLinkedListIsEmpty
	}

	s.Count--

	if s.IsEmpty() {
		s.Head = nil
		s.Tail = nil
	} else {
		current := s.Head
		for current.Next != s.Tail {
			current = current.Next
		}

		current.Next = nil
		s.Tail = current
	}

	return nil
}

func (s *SinglyLinkedList) IsEmpty() bool {
	return s.Count == 0
}

func (s *SinglyLinkedList) Size() int {
	return s.Count
}

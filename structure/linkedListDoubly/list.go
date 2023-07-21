package linkedListDoubly

import "errors"

var ErrDoublyLinkedListIsEmpty = errors.New("doubly linked list is empty")

type DoublyLinkedList struct {
	Count int
	Head  *Node
	Tail  *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (s *DoublyLinkedList) AddFirst(val int) {
	node := NewNode(val)

	if s.IsEmpty() {
		s.Head = node
		s.Tail = node
	} else {
		s.Head.Prev = node
		node.Next = s.Head
		s.Head = node
	}

	s.Head.Prev = s.Tail
	s.Tail.Next = s.Head

	s.Count++
}

func (s *DoublyLinkedList) AddLast(val int) {
	node := NewNode(val)

	if s.IsEmpty() {
		s.Head = node
		s.Tail = node
	} else {
		node.Prev = s.Tail
		s.Tail.Next = node
		s.Tail = node
	}

	s.Head.Prev = s.Tail
	s.Tail.Next = s.Head

	s.Count++
}

func (s *DoublyLinkedList) RemoveFirst() error {
	if s.IsEmpty() {
		return ErrDoublyLinkedListIsEmpty
	}
	s.Count--

	if s.IsEmpty() {
		s.Head = nil
		s.Tail = nil
	} else {
		s.Head = s.Head.Next

		s.Head.Prev = s.Tail
		s.Tail.Next = s.Head
	}

	return nil
}

func (s *DoublyLinkedList) RemoveLast() error {
	if s.IsEmpty() {
		return ErrDoublyLinkedListIsEmpty
	}

	s.Count--

	if s.IsEmpty() {
		s.Head = nil
		s.Tail = nil
	} else {
		s.Tail = s.Tail.Prev

		s.Head.Prev = s.Tail
		s.Tail.Next = s.Head
	}

	return nil
}

func (s *DoublyLinkedList) IsEmpty() bool {
	return s.Count == 0
}

func (s *DoublyLinkedList) Size() int {
	return s.Count
}

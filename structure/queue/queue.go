package queue

import (
	"errors"
	"fmt"
)

var ErrQueueIsEmpty = errors.New("queue is empty")

type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(val int) {
	q.items = append(q.items, val)

	fmt.Printf("data: %v, len: %d, cap: %d \n\r", q.items, len(q.items), cap(q.items))
}

func (q *Queue) Pop() (int, error) {
	if q.IsEmpty() {
		return 0, ErrQueueIsEmpty
	}

	val := q.items[0]
	q.items = q.items[1:len(q.items)]

	fmt.Printf("data: %v, len: %d, cap: %d \n\r", q.items, len(q.items), cap(q.items))

	return val, nil
}

func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, ErrQueueIsEmpty
	}

	return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Render() {
	fmt.Println(q.items)
}

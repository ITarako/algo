package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_New1(t *testing.T) {
	queue := NewQueue()

	assert.Equal(t, 0, queue.Size())
	assert.True(t, queue.IsEmpty())
}

func TestDoublyLinkedList_Push(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		queue := NewQueue()
		queue.Push(1)

		assert.Equal(t, 1, queue.Size())
		assert.False(t, queue.IsEmpty())
	})

	t.Run("2", func(t *testing.T) {
		queue := NewQueue()
		queue.Push(1)
		queue.Push(2)
		queue.Push(3)

		assert.Equal(t, 3, queue.Size())
		assert.False(t, queue.IsEmpty())
	})
}

func TestDoublyLinkedList_Pop(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		queue := NewQueue()

		_, err := queue.Pop()

		assert.EqualError(t, err, ErrQueueIsEmpty.Error())
	})

	t.Run("2", func(t *testing.T) {
		queue := NewQueue()
		queue.Push(1)
		val, err := queue.Pop()

		assert.Equal(t, 1, val)
		assert.Equal(t, 0, queue.Size())
		assert.True(t, queue.IsEmpty())
		assert.Nil(t, err)
	})

	t.Run("3", func(t *testing.T) {
		queue := NewQueue()
		queue.Push(1)
		queue.Push(2)
		val, err := queue.Pop()

		assert.Equal(t, 1, val)
		assert.Equal(t, 1, queue.Size())
		assert.False(t, queue.IsEmpty())
		assert.Nil(t, err)

		val, err = queue.Pop()

		assert.Equal(t, 2, val)
		assert.Equal(t, 0, queue.Size())
		assert.True(t, queue.IsEmpty())
		assert.Nil(t, err)
	})

	t.Run("4", func(t *testing.T) {
		queue := NewQueue()
		queue.Push(1)
		queue.Push(2)
		queue.Push(3)
		queue.Push(4)
		val, err := queue.Pop()

		assert.Equal(t, 1, val)
		assert.Equal(t, 3, queue.Size())
		assert.False(t, queue.IsEmpty())
		assert.Nil(t, err)

		val, err = queue.Pop()
		val, err = queue.Pop()

		assert.Equal(t, 3, val)
		assert.Equal(t, 1, queue.Size())
		assert.False(t, queue.IsEmpty())
		assert.Nil(t, err)

		val, err = queue.Pop()

		assert.Equal(t, 4, val)
		assert.Equal(t, 0, queue.Size())
		assert.True(t, queue.IsEmpty())
		assert.Nil(t, err)
	})
}

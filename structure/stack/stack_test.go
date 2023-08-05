package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_New1(t *testing.T) {
	stack := NewStack()

	assert.Equal(t, 0, stack.Size())
	assert.True(t, stack.IsEmpty())
}

func TestDoublyLinkedList_Push(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		stack := NewStack()
		stack.Push(1)

		assert.Equal(t, 1, stack.Size())
		assert.False(t, stack.IsEmpty())
	})

	t.Run("2", func(t *testing.T) {
		stack := NewStack()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		assert.Equal(t, 3, stack.Size())
		assert.False(t, stack.IsEmpty())
	})
}

func TestDoublyLinkedList_Pop(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		stack := NewStack()

		_, err := stack.Pop()

		assert.EqualError(t, err, ErrStackIsEmpty.Error())
	})

	t.Run("2", func(t *testing.T) {
		stack := NewStack()
		stack.Push(1)
		val, err := stack.Pop()

		assert.Equal(t, 1, val)
		assert.Equal(t, 0, stack.Size())
		assert.True(t, stack.IsEmpty())
		assert.Nil(t, err)
	})

	t.Run("3", func(t *testing.T) {
		stack := NewStack()
		stack.Push(1)
		stack.Push(2)
		val, err := stack.Pop()

		assert.Equal(t, 2, val)
		assert.Equal(t, 1, stack.Size())
		assert.False(t, stack.IsEmpty())
		assert.Nil(t, err)

		val, err = stack.Pop()

		assert.Equal(t, 1, val)
		assert.Equal(t, 0, stack.Size())
		assert.True(t, stack.IsEmpty())
		assert.Nil(t, err)
	})

	t.Run("4", func(t *testing.T) {
		stack := NewStack()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		stack.Push(4)
		val, err := stack.Pop()

		assert.Equal(t, 4, val)
		assert.Equal(t, 3, stack.Size())
		assert.False(t, stack.IsEmpty())
		assert.Nil(t, err)

		val, err = stack.Pop()
		val, err = stack.Pop()

		assert.Equal(t, 2, val)
		assert.Equal(t, 1, stack.Size())
		assert.False(t, stack.IsEmpty())
		assert.Nil(t, err)

		val, err = stack.Pop()

		assert.Equal(t, 1, val)
		assert.Equal(t, 0, stack.Size())
		assert.True(t, stack.IsEmpty())
		assert.Nil(t, err)
	})
}

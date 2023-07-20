package linkedListSingly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSinglyLinkedList_New1(t *testing.T) {
	list := NewSinglyLinkedList()

	assert.Nil(t, list.Head)
	assert.Nil(t, list.Tail)
	assert.Equal(t, 0, list.Size())
	assert.True(t, list.IsEmpty())
}

func TestSinglyLinkedList_AddFirst(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddFirst(1)

		assert.Equal(t, 1, list.Head.Val)
		assert.Equal(t, 1, list.Tail.Val)
		assert.Equal(t, list.Head.Val, list.Tail.Val)
		assert.Equal(t, 1, list.Size())
		assert.False(t, list.IsEmpty())
	})

	t.Run("2", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)
		list.AddFirst(4)

		assert.Equal(t, 4, list.Head.Val)
		assert.Equal(t, 1, list.Tail.Val)
		assert.Equal(t, 4, list.Size())
		assert.False(t, list.IsEmpty())
	})
}

func TestSinglyLinkedList_AddLast(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddLast(1)

		assert.Equal(t, 1, list.Head.Val)
		assert.Equal(t, 1, list.Tail.Val)
		assert.Equal(t, list.Head.Val, list.Tail.Val)
		assert.Equal(t, 1, list.Size())
		assert.False(t, list.IsEmpty())
	})

	t.Run("2", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddLast(1)
		list.AddLast(2)
		list.AddLast(3)
		list.AddLast(4)

		assert.Equal(t, 1, list.Head.Val)
		assert.Equal(t, 4, list.Tail.Val)
		assert.Equal(t, 4, list.Size())
		assert.False(t, list.IsEmpty())
	})
}

func TestSinglyLinkedList_AddFirstRemoveFirst(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		list := NewSinglyLinkedList()
		err := list.RemoveFirst()

		assert.EqualError(t, err, ErrSinglyLinkedListIsEmpty.Error())
	})

	t.Run("2", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddFirst(1)
		err := list.RemoveFirst()

		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.IsEmpty())
		assert.Nil(t, err)
	})

	t.Run("3", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddFirst(1)
		list.AddFirst(2)
		err := list.RemoveFirst()

		assert.Equal(t, 1, list.Head.Val)
		assert.Equal(t, 1, list.Tail.Val)
		assert.Equal(t, list.Head.Val, list.Tail.Val)
		assert.Equal(t, 1, list.Size())
		assert.False(t, list.IsEmpty())
		assert.Nil(t, err)

		err = list.RemoveFirst()

		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.IsEmpty())
		assert.Nil(t, err)
	})

	t.Run("4", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)
		list.AddFirst(4)
		err := list.RemoveFirst()

		assert.Equal(t, 3, list.Head.Val)
		assert.Equal(t, 1, list.Tail.Val)
		assert.Equal(t, 3, list.Size())
		assert.False(t, list.IsEmpty())
		assert.Nil(t, err)

		err = list.RemoveFirst()
		err = list.RemoveFirst()

		assert.Equal(t, 1, list.Head.Val)
		assert.Equal(t, 1, list.Tail.Val)
		assert.Equal(t, list.Head.Val, list.Tail.Val)
		assert.Equal(t, 1, list.Size())
		assert.False(t, list.IsEmpty())
		assert.Nil(t, err)

		err = list.RemoveFirst()

		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.IsEmpty())
		assert.Nil(t, err)
	})
}

func TestSinglyLinkedList_AddFirstRemoveLast(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		list := NewSinglyLinkedList()
		err := list.RemoveLast()

		assert.EqualError(t, err, ErrSinglyLinkedListIsEmpty.Error())
	})

	t.Run("2", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddFirst(1)
		err := list.RemoveLast()

		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.IsEmpty())
		assert.Nil(t, err)
	})

	t.Run("3", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddFirst(1)
		list.AddFirst(2)
		err := list.RemoveLast()

		assert.Equal(t, 2, list.Head.Val)
		assert.Equal(t, 2, list.Tail.Val)
		assert.Equal(t, list.Head.Val, list.Tail.Val)
		assert.Equal(t, 1, list.Size())
		assert.False(t, list.IsEmpty())
		assert.Nil(t, err)

		err = list.RemoveLast()

		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.IsEmpty())
		assert.Nil(t, err)
	})

	t.Run("4", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)
		list.AddFirst(4)
		err := list.RemoveLast()

		assert.Equal(t, 4, list.Head.Val)
		assert.Equal(t, 2, list.Tail.Val)
		assert.Equal(t, 3, list.Size())
		assert.False(t, list.IsEmpty())
		assert.Nil(t, err)

		err = list.RemoveLast()
		err = list.RemoveLast()

		assert.Equal(t, 4, list.Head.Val)
		assert.Equal(t, 4, list.Tail.Val)
		assert.Equal(t, list.Head.Val, list.Tail.Val)
		assert.Equal(t, 1, list.Size())
		assert.False(t, list.IsEmpty())
		assert.Nil(t, err)

		err = list.RemoveLast()

		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.IsEmpty())
		assert.Nil(t, err)
	})
}

func TestSinglyLinkedList_AddLastRemoveFirst(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		list := NewSinglyLinkedList()
		err := list.RemoveFirst()

		assert.EqualError(t, err, ErrSinglyLinkedListIsEmpty.Error())
	})

	t.Run("2", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddLast(1)
		err := list.RemoveFirst()

		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.IsEmpty())
		assert.Nil(t, err)
	})

	t.Run("3", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddLast(1)
		list.AddLast(2)
		err := list.RemoveFirst()

		assert.Equal(t, 2, list.Head.Val)
		assert.Equal(t, 2, list.Tail.Val)
		assert.Equal(t, list.Head.Val, list.Tail.Val)
		assert.Equal(t, 1, list.Size())
		assert.False(t, list.IsEmpty())
		assert.Nil(t, err)

		err = list.RemoveFirst()

		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.IsEmpty())
		assert.Nil(t, err)
	})

	t.Run("4", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddLast(1)
		list.AddLast(2)
		list.AddLast(3)
		list.AddLast(4)
		err := list.RemoveFirst()

		assert.Equal(t, 2, list.Head.Val)
		assert.Equal(t, 4, list.Tail.Val)
		assert.Equal(t, 3, list.Size())
		assert.False(t, list.IsEmpty())
		assert.Nil(t, err)

		err = list.RemoveFirst()
		err = list.RemoveFirst()

		assert.Equal(t, 4, list.Head.Val)
		assert.Equal(t, 4, list.Tail.Val)
		assert.Equal(t, list.Head.Val, list.Tail.Val)
		assert.Equal(t, 1, list.Size())
		assert.False(t, list.IsEmpty())
		assert.Nil(t, err)

		err = list.RemoveFirst()

		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.IsEmpty())
		assert.Nil(t, err)
	})
}

func TestSinglyLinkedList_AddLastRemoveLast(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		list := NewSinglyLinkedList()
		err := list.RemoveLast()

		assert.EqualError(t, err, ErrSinglyLinkedListIsEmpty.Error())
	})

	t.Run("2", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddLast(1)
		err := list.RemoveLast()

		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.IsEmpty())
		assert.Nil(t, err)
	})

	t.Run("3", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddLast(1)
		list.AddLast(2)
		err := list.RemoveLast()

		assert.Equal(t, 1, list.Head.Val)
		assert.Equal(t, 1, list.Tail.Val)
		assert.Equal(t, list.Head.Val, list.Tail.Val)
		assert.Equal(t, 1, list.Size())
		assert.False(t, list.IsEmpty())
		assert.Nil(t, err)

		err = list.RemoveLast()

		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.IsEmpty())
		assert.Nil(t, err)
	})

	t.Run("4", func(t *testing.T) {
		list := NewSinglyLinkedList()
		list.AddLast(1)
		list.AddLast(2)
		list.AddLast(3)
		list.AddLast(4)
		err := list.RemoveLast()

		assert.Equal(t, 1, list.Head.Val)
		assert.Equal(t, 3, list.Tail.Val)
		assert.Equal(t, 3, list.Size())
		assert.False(t, list.IsEmpty())
		assert.Nil(t, err)

		err = list.RemoveLast()
		err = list.RemoveLast()

		assert.Equal(t, 1, list.Head.Val)
		assert.Equal(t, 1, list.Tail.Val)
		assert.Equal(t, list.Head.Val, list.Tail.Val)
		assert.Equal(t, 1, list.Size())
		assert.False(t, list.IsEmpty())
		assert.Nil(t, err)

		err = list.RemoveLast()

		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
		assert.Equal(t, 0, list.Size())
		assert.True(t, list.IsEmpty())
		assert.Nil(t, err)
	})
}

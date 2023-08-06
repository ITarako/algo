package sorting

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func FuzzBubbleSort(f *testing.F) {
	f.Add(0)
	f.Add(1)
	f.Add(5)
	f.Fuzz(func(t *testing.T, n int) {
		slice := generateIntSlice(n)
		r1 := make([]int, 0, len(slice))
		copy(r1, slice)
		r2 := make([]int, 0, len(slice))
		copy(r2, slice)

		BubbleSort(r1)
		sort.Ints(r2)

		assert.Equal(t, r1, r2)
	})
}

func FuzzSelectionSort(f *testing.F) {
	f.Add(0)
	f.Add(1)
	f.Add(5)
	f.Fuzz(func(t *testing.T, n int) {
		slice := generateIntSlice(n)
		r1 := make([]int, 0, len(slice))
		copy(r1, slice)
		r2 := make([]int, 0, len(slice))
		copy(r2, slice)

		SelectionSort(r1)
		sort.Ints(r2)

		assert.Equal(t, r1, r2)
	})
}

func FuzzInsertionSort(f *testing.F) {
	f.Add(0)
	f.Add(1)
	f.Add(5)
	f.Fuzz(func(t *testing.T, n int) {
		slice := generateIntSlice(n)
		r1 := make([]int, 0, len(slice))
		copy(r1, slice)
		r2 := make([]int, 0, len(slice))
		copy(r2, slice)

		InsertionSort(r1)
		sort.Ints(r2)

		assert.Equal(t, r1, r2)
	})
}

func FuzzShellSort(f *testing.F) {
	f.Add(0)
	f.Add(1)
	f.Add(5)
	f.Fuzz(func(t *testing.T, n int) {
		slice := generateIntSlice(n)
		r1 := make([]int, 0, len(slice))
		copy(r1, slice)
		r2 := make([]int, 0, len(slice))
		copy(r2, slice)

		ShellSort(r1)
		sort.Ints(r2)

		assert.Equal(t, r1, r2)
	})
}

func FuzzMergeSort(f *testing.F) {
	f.Add(0)
	f.Add(1)
	f.Add(5)
	f.Fuzz(func(t *testing.T, n int) {
		slice := generateIntSlice(n)
		r1 := make([]int, 0, len(slice))
		copy(r1, slice)
		r2 := make([]int, 0, len(slice))
		copy(r2, slice)

		MergeSort(r1)
		sort.Ints(r2)

		assert.Equal(t, r1, r2)
	})
}

func FuzzQuickSort(f *testing.F) {
	f.Add(0)
	f.Add(1)
	f.Add(5)
	f.Fuzz(func(t *testing.T, n int) {
		slice := generateIntSlice(n)
		r1 := make([]int, 0, len(slice))
		copy(r1, slice)
		r2 := make([]int, 0, len(slice))
		copy(r2, slice)

		QuickSort(r1)
		sort.Ints(r2)

		assert.Equal(t, r1, r2)
	})
}

func generateIntSlice(n int) []int {
	n %= 20
	var values []int
	for i := 0; i < n; i++ {
		val := rand.Int() % 1e6
		values = append(values, val)
	}

	return values
}

package sorting

import "testing"

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort([]int{5, 1, 7, 2, 9, 9, 3})
	}
}

func BenchmarkBubbleSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort2([]int{5, 1, 7, 2, 9, 9, 3})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort([]int{5, 1, 7, 2, 9, 9, 3})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort([]int{5, 1, 7, 2, 9, 9, 3})
	}
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShellSort([]int{5, 1, 7, 2, 9, 9, 3})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort([]int{5, 1, 7, 2, 9, 9, 3})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort([]int{5, 1, 7, 2, 9, 9, 3})
	}
}

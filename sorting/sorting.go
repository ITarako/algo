package sorting

func BubbleSort2(input []int) {
	for i := 0; i < len(input); i++ {
		ii := i
		for j := 0; j < len(input); j++ {
			if input[ii] > input[j] && ii < j {
				swap(input, ii, j)
				ii = j
			}
		}
	}
}

func BubbleSort(input []int) {
	for wall := len(input) - 1; wall > 0; wall-- {
		for i := 0; i < wall; i++ {
			if input[i] > input[i+1] {
				swap(input, i, i+1)
			}
		}
	}
}

func SelectionSort(input []int) {
	for wall := len(input) - 1; wall > 0; wall-- {
		largestIdx := 0
		for i := 1; i <= wall; i++ {
			if input[i] > input[largestIdx] {
				largestIdx = i
			}
		}

		swap(input, largestIdx, wall)
	}
}

func InsertionSort(input []int) {
	for wall := 1; wall < len(input); wall++ {
		curUnsorted := input[wall]
		var i int
		for i = wall; i > 0; i-- {
			if curUnsorted >= input[i-1] {
				break
			}
			input[i] = input[i-1]
		}
		input[i] = curUnsorted
	}
}

func ShellSort(input []int) {
	gap := 1
	for gap < len(input)/3 {
		gap = 3*gap + 1
	}

	for gap >= 1 {
		for i := gap; i < len(input); i++ {
			for j := i; j >= gap && input[j] < input[j-gap]; j -= gap {
				swap(input, j, j-gap)
			}
		}

		gap /= 3
	}
}

func swap(input []int, i, j int) {
	if i == j {
		return
	}

	tmp := input[i]
	input[i] = input[j]
	input[j] = tmp
}

package sorting

func BubbleSort(input []int) []int {
	for partIndex := len(input) - 1; partIndex > 0; partIndex-- {
		for i := 0; i < partIndex; i++ {
			if input[i] > input[i+1] {
				swap(input, i, i+1)
			}
		}
	}

	return input
}

func SelectionSort(input []int) []int {
	for partIndex := len(input) - 1; partIndex > 0; partIndex-- {
		var largestAt = 0
		for i := 1; i <= partIndex; i++ {
			if input[i] > input[largestAt] {
				largestAt = i
			}
		}
		swap(input, largestAt, partIndex)
	}

	return input
}

func InsertionSort(input []int) {
	for partIndex := 1; partIndex < len(input); partIndex++ {
		curUnsorted := input[partIndex]
		var i int
		for i = partIndex; i > 0 && input[i-1] > curUnsorted; i-- {
			input[i] = input[i-1]
		}
		input[i] = curUnsorted
	}
}

func ShellSort(input []int) {
	var gap int = 1
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

	temp := input[i]
	input[i] = input[j]
	input[j] = temp
}

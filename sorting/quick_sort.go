package sorting

type quickSorter struct {
	input []int
}

func QuickSort(input []int) {
	sorter := quickSorter{
		input: input,
	}

	sorter.sort(0, len(input)-1)
}

func (s quickSorter) sort(low, high int) {
	if high <= low {
		return
	}

	j := s.partition(low, high)

	s.sort(low, j-1)
	s.sort(j+1, high)
}

func (s quickSorter) partition(low, high int) int {
	i := low
	j := high + 1

	pivot := s.input[low]

	for {
		i++
		for s.input[i] < pivot {
			if i == high {
				break
			}
			i++
		}

		j--
		for pivot < s.input[j] {
			if j == low {
				break
			}
			j--
		}

		if i >= j {
			break
		}

		swap(s.input, i, j)
	}

	swap(s.input, low, j)

	return j
}

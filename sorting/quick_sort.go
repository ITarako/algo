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

	pivot := s.partition(low, high)

	s.sort(low, pivot-1)
	s.sort(pivot+1, high)
}

func (s quickSorter) partition(low, high int) int {
	//random := rand.Int()
	//pivot := random%(high-low+1) + low
	//swap(s.input, pivot, low)

	pivot := low
	i := low
	j := low

	pivotVal := s.input[low]

	for k := low + 1; k <= high; k++ {
		if s.input[k] >= pivotVal {
			j++
		} else {
			i++
			swap(s.input, i, k)
			j++
		}
	}

	swap(s.input, pivot, i)

	return i
}

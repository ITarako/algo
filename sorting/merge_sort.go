package sorting

type mergeSorter struct {
	input []int
	aux   []int
}

func MergeSort(input []int) {
	sorter := mergeSorter{
		input: input,
		aux:   make([]int, len(input)),
	}

	sorter.sort(0, len(input)-1)
}

func (s mergeSorter) sort(low, high int) {
	if high <= low {
		return
	}

	mid := (low + high) / 2

	s.sort(low, mid)
	s.sort(mid+1, high)

	s.merge(low, mid, high)
}

func (s mergeSorter) merge(low, mid, high int) {
	s.fillAux(low, high)

	i := low
	j := mid + 1

	for k := low; k <= high; k++ {
		if i > mid {
			s.input[k] = s.aux[j]
			j++
		} else if j > high {
			s.input[k] = s.aux[i]
			i++
		} else if s.aux[j] < s.aux[i] {
			s.input[k] = s.aux[j]
			j++
		} else {
			s.input[k] = s.aux[i]
			i++
		}
	}
}

func (s mergeSorter) fillAux(low, high int) {
	for i := low; i <= high; i++ {
		s.aux[i] = s.input[i]
	}
}

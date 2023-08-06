package search

import "errors"

var ErrSearchNotFound = errors.New("value not found")

func Binary(val int, input []int) (int, error) {
	low := 0
	high := len(input) - 1

	for low <= high {
		mid := (high + low) / 2

		if input[mid] == val {
			return mid, nil
		} else if val > input[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return 0, ErrSearchNotFound
}

func BinaryRecursive(val int, input []int) (int, error) {
	return binaryRecursive(0, len(input)-1, val, input)
}

func binaryRecursive(low, high, val int, input []int) (int, error) {
	if low > high {
		return 0, ErrSearchNotFound
	}

	mid := (high + low) / 2
	if input[mid] == val {
		return mid, nil
	} else if val > input[mid] {
		return binaryRecursive(mid+1, high, val, input)
	} else {
		return binaryRecursive(low, mid-1, val, input)
	}
}

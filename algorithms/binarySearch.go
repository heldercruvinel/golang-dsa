package algorithms

import (
	"errors"
)

// Binary search requires a sorted list containing elements of the same type
func binarySearch(toFind int, list []int) (result int, err error) {

	length := len(list)

	if length == 1 {
		if list[0] == toFind {
			return list[0], nil
		}

		return 0, errors.New("elemento não encontrado")
	}

	half := length / 2
	first := list[:half]
	second := list[half:length]

	if toFind <= first[len(first)-1] {
		return binarySearch(toFind, first)
	}

	return binarySearch(toFind, second)
}

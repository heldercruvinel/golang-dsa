package algorithms

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	t.Run("Should to find element 3 on a integer list", func(t *testing.T) {
		toFind := 3
		intList := make([]int, 0, 5)
		intList = append(intList, 1, 2, 3, 4, 5)

		result, err := binarySearch(toFind, intList)
		if err != nil {
			t.Errorf("Failed, want: %d, got: %d", toFind, result)
		}

	})

	t.Run("Should to find element 6 on a integer list", func(t *testing.T) {
		toFind := 6
		intList := make([]int, 0, 5)
		intList = append(intList, 1, 2, 3, 4, 5, 6)

		result, err := binarySearch(toFind, intList)
		if err != nil {
			t.Errorf("Failed, want: %d, got: %d, err: %q", toFind, result, err.Error())
		}

	})

	t.Run("Should NOT to find element 7 on a integer list", func(t *testing.T) {
		toFind := 7
		intList := make([]int, 0, 5)
		intList = append(intList, 1, 2, 3, 4, 5)

		_, err := binarySearch(toFind, intList)
		if err.Error() != "elemento não encontrado" {
			t.Errorf("Failed, want: elemento não encontrado, got: %q", err.Error())
		}

	})

	t.Run("Should to find element 900000 on a big integer list", func(t *testing.T) {
		toFind := 999999999
		intList := make([]int, 1000000000)

		for i := 1; i <= 1000000; i++ {
			intList[i-1] = i
		}

		result, err := binarySearch(toFind, intList)
		fmt.Println(result)
		if result != toFind && err == nil {
			t.Errorf("Failed, want: %d, got: %d", toFind, result)
		}

	})

}

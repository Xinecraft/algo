package searching

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	sortedArray := []int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9}
	needle := 9
	result := BinarySearch(sortedArray, needle)
	if result != 9 {
		t.Errorf("Error")
	}
}

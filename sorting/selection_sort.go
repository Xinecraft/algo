package sorting

func SelectionSort[T Number](array []T) {
	for i := 0; i < len(array); i++ {
		min_idx := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min_idx] {
				min_idx = j
			}
		}
		SwapItem(i, min_idx, array)
	}
}

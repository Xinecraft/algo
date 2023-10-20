package sorting

func SelectionSort[T Number](array []T) {
	for i := 0; i < len(array); i++ {
		minIdx := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}
		SwapItem(i, minIdx, array)
	}
}

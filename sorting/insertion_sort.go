package sorting

func InsertionSort[T Number](array []T) {
	for i := 1; i < len(array); i++ {
		currIdx := i
		for currIdx > 0 && array[currIdx-1] > array[currIdx] {
			SwapItem(currIdx-1, currIdx, array)
			currIdx--
		}
	}
}

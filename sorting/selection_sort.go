package sorting

type Number interface {
	int | float32 | float64
}

func SelectionSort[T Number](array []T) {
	for i := 0; i < len(array); i++ {
		min_idx := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min_idx] {
				min_idx = j
			}
		}
		swapItem(i, min_idx, array)
	}
}

func swapItem[T Number](i int, j int, array []T) {
	var temp T = array[i]
	array[i] = array[j]
	array[j] = temp
}

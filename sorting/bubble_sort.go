package sorting

func BubbleSort[T Number](array []T) {
	hasSwapped := true
	for hasSwapped == true {
		hasSwapped = false
		for i := 0; i < len(array)-1-i; i++ {
			if array[i] > array[i+1] {
				SwapItem(i, i+1, array)
				hasSwapped = true
			}
		}
	}
}

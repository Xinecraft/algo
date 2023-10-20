package sorting

func BubbleSort[T Number](array []T) {
	hasSwapped := true
	k := 0
	for hasSwapped == true {
		hasSwapped = false
		for i := 0; i < len(array)-1-k; i++ {
			if array[i] > array[i+1] {
				hasSwapped = true
				SwapItem(i, i+1, array)
			}
		}
		k++
	}
}

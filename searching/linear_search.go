package searching

func LinearSearch[T int | float32 | float64](array []T, needle T) int {
	for i := 0; i < len(array); i++ {
		if array[i] == needle {
			return i
		}
	}
	return -1
}

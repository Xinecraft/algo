package searching

func BinarySearch[T int | float32 | float64](array []T, needle T) int {
	start := 0
	end := len(array) - 1

	for start <= end {
		mid := (start + end) / 2

		if array[mid] == needle {
			return mid
		} else if array[mid] < needle {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

package sorting

type Number interface {
	int | float32 | float64
}

func SwapItem[T Number](i int, j int, array []T) {
	var temp T = array[i]
	array[i] = array[j]
	array[j] = temp
}

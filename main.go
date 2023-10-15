package main

import (
	"fmt"
	"github.com/xinecraft/algo/searching"
	"github.com/xinecraft/algo/sorting"
)

func main() {
	fmt.Println("Hello World")

	arr := []int{1, 2, 5, 3, 1, 4, 6, 8}
	fmt.Println(arr)
	idx := searching.LinearSearch(arr, 4)
	//notfound := searching.LinearSearch(arr, 999)
	fmt.Println(idx)
	//fmt.Println(notfound)

	sorting.SelectionSort(arr)
	fmt.Print(arr)
}

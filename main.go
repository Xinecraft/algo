package main

import (
	//"github.com/xinecraft/algo/searching"
	// "github.com/xinecraft/algo/sorting"
	"github.com/xinecraft/algo/ds"
)

func main() {
	//arr := []int{1, 2, 5, 3, 1, 4, 6, 8}
	// [1,1,4,2,1,3]
	// arr2 := []int{1, 1, 4, 2, 1, 3}
	//fmt.Println(arr)
	// fmt.Println(arr2)
	//idx := searching.LinearSearch(arr, 4)
	//notfound := searching.LinearSearch(arr, 999)
	//fmt.Println(idx)
	//fmt.Println(notfound)

	//sorting.SelectionSort(arr)
	//fmt.Println(arr)

	//sorting.BubbleSort(arr2)
  // sorting.InsertionSort(arr2)
  // fmt.Println(arr2)

  list := ds.LinkedList[string]{}
  list.Append("1")
  list.Append("2")
  list.Append("3")
  list.Prepend("sda")
  list.Print()
  print("\nLength", list.Length)
}

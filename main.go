package main

import (
	"fmt"

	"github.com/xinecraft/algo/dynamic_connectivity"
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

	// LinkedList
	// list := ds.LinkedList[string]{}
	// list.Append("1")
	// list.Append("2")
	// list.Append("3")
	// list.Prepend("sda")
	// fmt.Println(list)
	// fmt.Println("Length", list.Length)
	// fmt.Println("Find:", list.Find("sda"))
	// value, ok := list.GetAt(3)
	// fmt.Println("Get:", value, "Ok:", ok)
	// fmt.Println("Removing")
	// list.RemoveAt(0)
	// fmt.Println(list)
	// fmt.Println("Inserting")
	// list.InsertAt("kakamora", 2)
	// fmt.Println(list)

	// Queue
	// queue := ds.Queue[int]{}
	// queue.Enqueue(1)
	// queue.Enqueue(7)
	// queue.Enqueue(5)
	// fmt.Println(queue.Peek())
	// item := queue.Dequeue()
	// fmt.Println(item)
	// fmt.Println(queue.Peek())
	// fmt.Println(queue.Size())

	// Stack
	// stack := ds.Stack[float32]{}
	// stack.Push(3.10)
	// stack.Push(8.50)
	// stack.Push(11.50)
	// stack.Push(17.50)
	// fmt.Printf("stack: %v\n", stack)
	// stack.Push(19.50)
	// fmt.Printf("stack: %v\n", stack)
	// fmt.Println(stack.Pop())
	// fmt.Printf("stack: %v\n", stack)
	// stack.Push(33.50)
	// fmt.Printf("stack: %v\n", stack)
	// fmt.Println(stack.Pop())
	// fmt.Printf("stack: %v\n", stack)

	qf := dynamic_connectivity.MakeQuickFind(10)
	qf.Union(0, 1)
	qf.Union(1, 9)
	qf.Union(5, 6)
	qf.Union(5, 1)
	fmt.Println("Connected: ", qf.Connected(5, 9))

	qu := dynamic_connectivity.MakeQuickUnion(10)
	qu.Union(0, 1)
	qu.Union(1, 9)
	qu.Union(5, 6)
	qu.Union(5, 1)
	fmt.Println("Connected: ", qu.Connected(5, 9))

	quw := dynamic_connectivity.MakeQuickUnionWeighted(10)
	quw.Union(0, 1)
	quw.Union(1, 9)
	quw.Union(5, 6)
	quw.Union(5, 1)
	fmt.Println("Connected: ", quw.Connected(5, 9))

	quwpc := dynamic_connectivity.MakeQuickUnionWeightedPathCompressed(10)
	quwpc.Union(0, 1)
	quwpc.Union(1, 9)
	quwpc.Union(5, 6)
	quwpc.Union(5, 1)
	fmt.Println("Connected: ", quwpc.Connected(5, 9))
}

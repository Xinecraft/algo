package ds

import (
	"fmt"
	"strings"
)

type Node[T comparable] struct {
  value T
  next *Node[T]
  prev *Node[T]
}

type LinkedList[T comparable] struct {
  head *Node[T]
  tail *Node[T]
  Length int
}

/**
Add Item to the End of LinkedList
*/
func (l *LinkedList[T]) Append(item T) bool {
  l.Length = l.Length + 1

  newNode := Node[T]{value: item}
  newNode.prev = l.tail

  if (l.head == nil) {
    l.head = &newNode
  }

  if (l.tail != nil) {
    l.tail.next = &newNode
  }

  l.tail = &newNode

  return true
}

/**
Add Item to the Start of LinkedList
*/
func (l *LinkedList[T]) Prepend(item T) bool {
  l.Length = l.Length + 1

  newNode := Node[T]{value: item}
  newNode.next = l.head

  if (l.tail == nil) {
    l.tail = &newNode
  }

  if (l.head != nil) {
    l.head.prev = &newNode
  }

  l.head = &newNode

  return true
}

func (l *LinkedList[T]) InsertAt(item T, idx int) bool  {
  curr := l.getNode(idx)

  if curr == nil {
    return false
  }

  newNode := Node[T]{value: item}
  newNode.prev = curr.prev
  newNode.next = curr

  if curr.prev  != nil {
    curr.prev.next = &newNode
  }

  if idx == 0 {
    l.head = curr
  } else if idx == l.Length - 1 {
    l.tail = curr
  }

  curr.prev = &newNode

  return true
}

/**
Remove the First Item of LinkedList
*/
func (l *LinkedList[T]) RemoveFirst() T  {
  return l.RemoveAt(0)
}

/**
Remove the Last Item of LinkedList
*/
func (l *LinkedList[T]) RemoveLast() T  {
  return l.RemoveAt(l.Length - 1)
}

/**
Search Item T and remove from LinkedList
*/
func (l *LinkedList[T]) Remove(item T) T  {
  _, idx := l.findNode(item)
  
  return l.RemoveAt(idx)
}

/**
Remove Nth Item from LinkedList
*/
func (l *LinkedList[T]) RemoveAt(idx int) T  {
  curr := l.getNode(idx)
  value := curr.value

  if curr == nil {
    return *new(T)
  }

  l.Length--

  if l.Length <= 0 {
    l.head = nil
    l.tail = nil
    return value
  }

  if (curr.prev != nil) {
    curr.prev.next = curr.next
  }

  if(curr.next != nil) {
    curr.next.prev = curr.prev
  }

  if idx == 0 {
    l.head = curr.next
  } else if idx == l.Length - 1 {
    l.tail = curr.prev
  }

  return value
}

/**
Get Nth Item from LinkedList
*/
func (l *LinkedList[T]) GetAt(idx int) (T, bool)  {
  node := l.getNode(idx)

  if node == nil {
    return *new(T),false
  }

  return node.value, true
}

/**
Find if Item T exists in LinkedList
*/
func (l *LinkedList[T]) Find(item T) int {
  _, idx := l.findNode(item)

  return idx
}

func (l LinkedList[T]) String() string {
  str := strings.Builder{}

  for i := l.head; i != nil; i = i.next {
    str.WriteString(fmt.Sprint(i.value))

    if(i.next != nil) {
      str.WriteString(" -> ")
    }
  }

  return str.String()
}


func (l *LinkedList[T]) getNode(idx int) *Node[T] {
  if idx > l.Length {
    return nil
  }

  currNode := l.head
  for i := 0; i < idx; i++ {
    currNode = currNode.next
  }

  return currNode
}

func (l *LinkedList[T]) findNode(item T) (*Node[T], int) {
  currNode := l.head
  idx := 0
  for currNode != nil {
    if currNode.value == item {
      return currNode, idx
    }
    currNode = currNode.next
    idx++
  }

  return nil, -1
}


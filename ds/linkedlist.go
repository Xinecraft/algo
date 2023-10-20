package ds

type Node[T any] struct {
  value T
  next *Node[T]
  prev *Node[T]
}

type LinkedList[T any] struct {
  head *Node[T]
  tail *Node[T]
  Length int
}

func (l *LinkedList[T]) Append(item T)  {
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
}

func (l *LinkedList[T]) Prepend(item T)  {
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
}

func (l *LinkedList[T]) Remove(item T)  {
  print("Removed")
}

func (l *LinkedList[T]) RemoveAt(idx int)  {
  if idx > l.Length {
    return
  }

  for i := 0; i < idx; i++ {
    print("Removed")
  }
}

func (l *LinkedList[T]) Get(idx int)  {
  print("get")
}

func (l *LinkedList[T]) Find(item T) {
  print("Finding")
}

func (l *LinkedList[T]) Print() {
  node := l.head
  for node != nil {
    print(" -> ", node.value)
    node = node.next
  }
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

func (l *LinkedList[T]) findNode(idx int) *Node[T] {
  currNode := l.head


  return currNode
}

package ds

import "fmt"

type Queue[T comparable] struct {
	queue  LinkedList[T]
	length int
}

func (q *Queue[T]) Enqueue(item T) {
	q.length++
	q.queue.Append(item)
}

func (q *Queue[T]) Dequeue() T {
	if q.length <= 0 {
		return *new(T)
	}

	q.length--
	value := q.queue.RemoveFirst()
	return value
}

func (q *Queue[T]) Peek() T {
	node := q.queue.head

	if node == nil {
		return *new(T)
	}

	return node.value
}

func (q *Queue[T]) Size() int {
	return q.length
}

func (q Queue[T]) String() string {
	str := fmt.Sprint(q.queue)
	return str
}

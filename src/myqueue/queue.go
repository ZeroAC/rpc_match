package myqueue

type Queue[T any] struct {
	elements []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (queue *Queue[T]) Enqueue(element T) {
	queue.elements = append(queue.elements, element)
}

func (queue *Queue[T]) Dequeue() T {
	element := queue.elements[0]
	queue.elements = queue.elements[1:len(queue.elements)]
	return element
}

func (queue *Queue[T]) Front() T {
	return queue.elements[0]
}

func (queue *Queue[T]) IsEmpty() bool {
	return len(queue.elements) == 0
}

func (queue *Queue[T]) Size() int {
	return len(queue.elements)
}

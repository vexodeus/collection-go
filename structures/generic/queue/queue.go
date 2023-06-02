package queue

type Queue[T any] struct {
	Slice []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}
func (q *Queue[T]) Len() int {
	return len(q.Slice)
}
func (q *Queue[T]) Peek() any {
	if q.IsEmpty() {
		return nil
	}
	return q.Slice[0]
}
func (q *Queue[T]) Enqueue(value T) {
	q.Slice = append(q.Slice, value)
}
func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		return *new(T)
	}
	res := q.Slice[0]
	if q.Len() == 1 {
		q.Slice = []T{}
		return res
	}
	q.Slice = q.Slice[1:]
	return res
}
func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}

package collection

type Queue struct {
	Slice []any
}

func NewQueue() *Queue {
	return &Queue{}
}
func (q *Queue) Len() int {
	return len(q.Slice)
}
func (q *Queue) Peek() any {
	return q.Slice[0]
}
func (q *Queue) Enqueue(value any) {
	q.Slice = append(q.Slice, value)
}
func (q *Queue) Dequeue() any {
	res := q.Slice[0]
	q.Slice = q.Slice[1:]
	return res
}
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

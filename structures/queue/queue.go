package queue

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
	if q.IsEmpty() {
		return nil
	}
	return q.Slice[0]
}
func (q *Queue) Enqueue(value any) {
	q.Slice = append(q.Slice, value)
}
func (q *Queue) Dequeue() any {
	if q.IsEmpty() {
		return nil
	}
	res := q.Slice[0]
	if q.Len() == 1 {
		q.Slice = []any{}
		return res
	}
	q.Slice = q.Slice[1:]
	return res
}
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

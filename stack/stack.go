package stack

type Stack struct {
	Slice []any
}

func NewStack() *Stack {
	return &Stack{}
}
func (s *Stack) Len() int {
	return len(s.Slice)
}
func (s *Stack) Peek() any {
	return s.Slice[:s.Len()]
}
func (s *Stack) Push(value any) {
	s.Slice = append(s.Slice, value)
}
func (s *Stack) Pop() any {
	res := s.Peek()
	s.Slice = s.Slice[:s.Len()]
	return res
}
func (s *Stack) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

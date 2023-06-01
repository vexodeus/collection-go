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
	if s.IsEmpty() {
		return nil
	}
	return s.Slice[:s.Len()]
}
func (s *Stack) Push(value any) {
	s.Slice = append(s.Slice, value)
}
func (s *Stack) Pop() any {
	if s.IsEmpty() {
		return nil
	}
	res := s.Peek()
	if s.Len() == 1 {
		s.Slice = []any{}
		return res
	}
	s.Slice = s.Slice[:s.Len()]
	return res
}
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

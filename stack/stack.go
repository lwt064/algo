package stack

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	return &Stack{data: make([]interface{}, 0)}
}

func (s *Stack) Push(d interface{}) {
	s.data = append(s.data, d)
}

func (s *Stack) Pop() {
	if len(s.data) > 0 {
		s.data = s.data[0 : len(s.data)-1]
	}
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Top() interface{} {
	if !s.Empty() {
		return s.data[len(s.data)-1]
	}
	return nil
}

package stacks

import "fmt"

type StackSliceImpl struct {
	stack []interface{}
}

func (s *StackSliceImpl) Init() {
	s.stack = make([]interface{}, 0)
}

func (s *StackSliceImpl) Push(i interface{}) {

	s.stack = append(s.stack, i)
}
func (s *StackSliceImpl) Pop() {
	s.stack = (s.stack)[:len(s.stack)-1]
}

func (s *StackSliceImpl) Peek() interface{} {
	fmt.Println(s.stack...)
	return (s.stack)[len(s.stack)-1]
}

func (s *StackSliceImpl) Size() int {
	return len(s.stack)
}

func (s *StackSliceImpl) IsEmpty() bool {
	return len(s.stack) == 0
}

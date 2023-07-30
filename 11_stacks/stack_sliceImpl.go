package stacks

import "fmt"

type sliceStack []interface{}

var StackSliceImpl Stack = &sliceStack{}

func (s *sliceStack) Init() {
	*s = make([]interface{}, 0)
}

func (s *sliceStack) Push(i interface{}) {

	*s = append(*s, i)
}
func (s *sliceStack) Pop() interface{} {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *sliceStack) Peek() interface{} {
	fmt.Println(*s...)
	return (*s)[len(*s)-1]
}

func (s *sliceStack) Size() int {
	return len(*s)
}

func (s *sliceStack) IsEmpty() bool {
	return len(*s) == 0
}

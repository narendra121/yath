package stacks

import (
	"math"
	linkedList "yath/10_linkedList"
)

type StackLLImpl struct {
	node *linkedList.Node
	size int
}

func (l *StackLLImpl) Init() {
	l.node = nil
	l.size = 0
}

func (l *StackLLImpl) Push(i interface{}) {
	tmp := linkedList.GetNode(i.(int))
	tmp.Next = l.node
	l.node = tmp
	l.size++
}

func (l *StackLLImpl) Pop() {
	if l.node == nil {
		return
	}
	tmp := l.node
	l.node = l.node.Next
	tmp.Next = nil
	tmp = nil
	l.size--
}
func (l *StackLLImpl) Size() int {
	return l.size
}
func (l *StackLLImpl) Peek() interface{} {
	if l.node == nil {
		return math.MaxInt
	}
	return l.node.Data
}

func (l *StackLLImpl) IsEmpty() bool {
	return l.size == 0
}

package stacks

import (
	"math"
	linkedList "yath/10_linkedList"
)

type linkedStack struct {
	node *linkedList.Node
	size int
}

var StackLinkedListImpl Stack = &linkedStack{}

func (l *linkedStack) Init() {
	l.node = nil
	l.size = 0
}

func (l *linkedStack) Push(i interface{}) {
	tmp := linkedList.GetNode(i.(int))
	tmp.Next = l.node
	l.node = tmp
	l.size++
}

func (l *linkedStack) Pop() interface{} {
	if l.node == nil {
		return math.MaxInt
	}
	res := l.node.Data
	tmp := l.node
	l.node = l.node.Next
	tmp.Next = nil
	tmp = nil
	l.size--
	return res
}
func (l *linkedStack) Size() int {
	return l.size
}
func (l *linkedStack) Peek() interface{} {
	if l.node == nil {
		return math.MaxInt
	}
	return l.node.Data
}

func (l *linkedStack) IsEmpty() bool {
	return l.size == 0
}

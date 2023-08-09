package queue

/*
Enque is adding into Queue
DeQueue is removing from Queue
GetFront get first Index
GetRear to have last index
*/
type Queue interface {
	Init(len, cap int)
	Size() int
	IsFull() bool
	IsEmpty() bool
	EnQueue(x interface{})
	DeQueue()
	GetFront() interface{}
	GetRear() interface{}
	GetQueue() interface{}
}

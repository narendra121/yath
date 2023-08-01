package dequeue

type DoublyEndedQueue interface {
	Init(cap int)
	Size() int
	IsFull() bool
	IsEmpty() bool
	InsertFront(x interface{})
	InsertRear(x interface{})
	DeleteFront()
	DeleteRear()
	GetFront() interface{}
	GetRear() interface{}
	GetArray() interface{}
}

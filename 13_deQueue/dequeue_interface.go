package dequeue

type DoublyEndedQueue interface {
	InsertFront(x interface{})
	InsertRear(x interface{})
	DeleteFromnt()
	DeleteRear()
	GetFront() interface{}
	GetRear() interface{}
	Isfull() bool
	IsEmpty() bool
	Size() int
}

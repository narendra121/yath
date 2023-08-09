package queue

import linkedList "yath/10_linkedList"

type QueueLinkedListImpl struct {
	front *linkedList.Node
	rear  *linkedList.Node
	size  int
}

func (q *QueueLinkedListImpl) Init(dummy1, dummy2 int) {
	q.front = &linkedList.Node{}
	q.rear = &linkedList.Node{}
	q.size = 0
}
func (q *QueueLinkedListImpl) EnQueue(x interface{}) {
	temp := linkedList.GetNode(x.(int))
	if q.front == nil {
		q.front, q.rear = temp, temp
		return
	}
	q.rear.Next = temp
	q.rear = temp
	q.size++
}
func (q *QueueLinkedListImpl) DeQueue() {
	if q.front == nil {
		return
	}
	q.front = q.front.Next
	if q.front == nil {
		q.rear = nil
	}
	q.size--
}
func (q *QueueLinkedListImpl) GetFront() interface{} {
	return q.front
}
func (q *QueueLinkedListImpl) GetRear() interface{} {
	return q.rear
}
func (q *QueueLinkedListImpl) IsEmpty() bool {
	return q.front == nil
}

func (q *QueueLinkedListImpl) IsFull() bool {
	//No need to implement
	return false
}
func (q *QueueLinkedListImpl) Size() int {
	return q.size
}
func (q *QueueLinkedListImpl) GetQueue() interface{} {
	//No need
	return nil
}

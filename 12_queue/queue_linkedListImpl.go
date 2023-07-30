package queue

import "yath/10_linkedList"

type queueLinkedList struct {
	front *linkedList.Node
	rear  *linkedList.Node
	size  int
}

var QueueLinkedListImpl Queue = &queueLinkedList{}

func (q *queueLinkedList) Init(dummy1, dummy2 int) {
	q.front = &linkedList.Node{}
	q.rear = &linkedList.Node{}
	q.size = 0
}
func (q *queueLinkedList) EnQueue(x int) {
	temp := linkedList.GetNode(x)
	if q.front == nil {
		q.front, q.rear = temp, temp
		return
	}
	q.rear.Next = temp
	q.rear = temp
	q.size++
}
func (q *queueLinkedList) DeQueue() {
	if q.front == nil {
		return
	}
	q.front = q.front.Next
	if q.front == nil {
		q.rear = nil
	}
	q.size--
}
func (q *queueLinkedList) GetFront() interface{} {
	return q.front
}
func (q *queueLinkedList) GetRear() interface{} {
	return q.rear
}
func (q *queueLinkedList) IsEmpty() bool {
	return q.front == nil
}

func (q *queueLinkedList) IsFull() bool {
	//No need to implement
	return false
}
func (q *queueLinkedList) Size() int {
	return q.size
}
func (q *queueLinkedList) GetQueue() interface{} {
	//No need
	return nil
}

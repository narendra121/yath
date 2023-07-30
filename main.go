package main

import "yath/queue"

func main() {
	q := queue.QueueArrImpl
	q.Init(10, 12)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	queue.ReverseQueueUsingStack(q)
}

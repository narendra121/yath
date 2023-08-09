package queue

import (
	stacks "yath/11_stacks"
)

func ReverseQueueUsingStack(queue Queue) Queue {
	stack := stacks.StackSliceImpl{}
	stack.Init()
	for !queue.IsEmpty() {
		stack.Push(queue.GetQueue().([]int)[queue.GetFront().(int)])
		queue.DeQueue()
	}
	for !stack.IsEmpty() {
		queue.EnQueue(stack.Peek().(int))
		stack.Pop()
	}
	return queue
}

func GenerateNumbersWithGivenDigits(q Queue, x, y, n int) {
	q.EnQueue(x)
	q.EnQueue(y)
	for i := 0; i < n; i++ {
		/*
		   curr =q.top()

		   q.pop()

		   q.push(curr+'5')
		   q.push(cur+'6')

		*/
	}
}

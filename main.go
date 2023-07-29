package main

import (
	"fmt"
	"yath/queue"
)

func main() {
	// normal.StartSnakesAndLaddersGame()
	// linkedList.TraverseCircularDoublyLinkedList(linkedList.CreateCircularDoublyLinkedList())
	// linkedList.TraverseSingleLinkedList(linkedList.ReverseSingleLinkedList(linkedList.CreateSingleLinkedList()))
	// var wg sync.WaitGroup
	// ch1 := make(chan int)
	// wg.Add(1)
	// go add(1, ch1, &wg)
	// fmt.Println(<-ch1)

	// wg.Wait()
	// linkedList.ReverseSLinkedListAtKGroups(linkedList.CreateSingleLinkedList(),3)

	// linkedList.TraverseSingleLinkedList(linkedList.SegregateEvenAndOddInSLinkedList(linkedList.CreateSingleLinkedList()))
	// m := normal.NewHashTable()
	// m.AddValueToSlice(1, 2, 3, 4, 2, 5, 1)
	// m.PrintAll()
	// fmt.Println(m.GetIndexOfValue(7))

	// linkedList.TraverseDoubleLinkedList(linkedList.CreateDoubleLinkedListFromArray([]int{1, 2, 3, 4}))
	// fmt.Println(arrays.PrintLeadersOfArray([]int{1, 2, 3, 4, 5}))

	// fmt.Println(arrays.MaximumDifferenceInArray([]int{1, 8, 3, 6, 5}))
	// fmt.Println(searching.GetThePeakElement([]int{5, 10, 12, 90, 18, 30, 76, 60}))
	// fmt.Println(searching.IsAnySumEqual([]int{1, 2, 5, 8}, 3))
	// fmt.Println(sorting.MergeSort([]int{3, 3, 6, 9}, []int{1, 2, 3, 4, 5, 6}))
	// m := []int{-12, 18, -10, 15, -12}
	// sorting.HeapSort([]int{10, 15, 50, 4, 20})
	// sorting.MeetingMaxGuests([]int{900, 600, 700}, []int{1000, 800, 730})
	// matrix.PrintSpiral([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}})
	// fmt.Println(matrix.SearchInRowWiseAndCloumnWiseSortedMatrix([][]int{{10, 20, 30, 40}, {15, 25, 35, 45}, {27, 29, 37, 48}, {32, 33, 39, 50}}, 39))
	// fmt.Println(string_prblms.LongestSistinctSubStr("abcadbd"))
	// var myStack stacks.LinkedStack
	// myStack.Init()
	// fmt.Println(myStack.Size())

	// myStack.Push(10)
	// myStack.Push(20)
	// myStack.IsEmpty()
	// fmt.Println(myStack.Peek())
	// fmt.Println(myStack.Pop())
	// fmt.Println(myStack.Peek())
	// fmt.Println(myStack.Size())

	// fmt.Println(leetcode.SingleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
	// stacks.StockSpan([]int{60, 10, 20, 15, 35, 50}, 6)

	queue.ArrQueue.Init(0, 5)
	queue.ArrQueue.EnQueue(1)
	queue.ArrQueue.EnQueue(2)
	queue.ArrQueue.EnQueue(3)
	queue.ArrQueue.EnQueue(4)
	queue.ArrQueue.EnQueue(4)
	queue.ArrQueue.EnQueue(4)
	fmt.Println(queue.ArrQueue) //
	queue.ArrQueue.DeQueue()
	fmt.Println(queue.ArrQueue) //
	queue.ArrQueue.DeQueue()
	fmt.Println(queue.ArrQueue) //
	queue.ArrQueue.EnQueue(5)
	queue.ArrQueue.DeQueue()
	fmt.Println(queue.ArrQueue) //
	fmt.Println()
	fmt.Println()
	fmt.Println()
	// var queueTemp1 queue.Queue
	// queueTemp1.Init(0, 5)
	// queueTemp1.RearEnQueue(1)
	// queueTemp1.RearEnQueue(2)
	// queueTemp1.RearEnQueue(3)
	// queueTemp1.RearEnQueue(4)
	// queueTemp1.RearEnQueue(4)
	// queueTemp1.RearEnQueue(4)
	// fmt.Println(queueTemp1)
	// queueTemp1.FrontDeQueue()
	// fmt.Println(queueTemp1)
	// queueTemp1.FrontDeQueue()
	// fmt.Println(queueTemp1)
	// queueTemp1.RearEnQueue(5)
	// queueTemp1.FrontDeQueue()
	// fmt.Println(queueTemp1)
}

// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M

// if sl[int(s[i])] > sl[int(s[i-1])] {
// 	res += sl[int(s[i])] - tmp
// 	tmp = 0
// } else if sl[int(s[i])] <= sl[int(s[i-1])] {
// 	res += sl[int(s[i])]
// } else {
// 	tmp += sl[int(s[i])]
// 	// res = res + (sl[int(s[i])])
// }
// tmp = 0

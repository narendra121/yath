package main

import (
	"fmt"
	"yath/searching"
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
	fmt.Println(searching.IsAnySumEqualTriplet([]int{1, 2, 5, 8}, 1))
}

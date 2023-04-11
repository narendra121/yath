package main

import (
	"sync"
	"yath/linkedList"
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
	linkedList.ReverseSLinkedListAtKGroups(linkedList.CreateSingleLinkedList(),3)

}
func add(a int, ch1 chan int, wg *sync.WaitGroup) {
	ch1 <- a + 1
	wg.Done()
}

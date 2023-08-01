package main

import (
	"fmt"
	dequeue "yath/13_deQueue"
)

func main() {
	// dequeue.MaximumsInallSubArraysOfSizeK([]int{20, 40, 30, 10, 60}, 3)
	m := dequeue.FirstCircularTourNaive([]int{4, 8, 7, 4}, []int{6, 5, 3, 5})
	fmt.Println(m)
}

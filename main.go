package main

import (
	"fmt"
	"sort"
)

func main() {
	a := removeElement([]int{3, 2, 2, 3}, 3)
	fmt.Println(a)
}
func removeElement(nums []int, val int) int {
	low, high := 0, len(nums)-1
	point := 0
	for low < high {
		if nums[low] == val {
			nums[low] = 0
			point += 1
		}
		if nums[high] == val {
			nums[high] = 0
			point += 1
		}
		low++
		high--
	}
	sort.Ints(nums)
	nums = nums[point:]
	fmt.Println(nums)
	fmt.Println(point)
	return len(nums)
}

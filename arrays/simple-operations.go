package arrays

import (
	"fmt"
	"math"
)

func SimpleSearchInArray(sl [5]int, search int) int {
	for _, val := range sl {
		if val == search {
			return 1
		}
	}
	return -1
}

func SimpleInsertInArray(array [5]int, capacity, len, newVal, position int) int {
	if position > capacity {
		return -1
	}
	idx := position - 1
	for i := len + 1; i > idx; i-- {
		array[i] = array[i-1]
	}
	array[idx] = newVal
	return 1
}

func SimpleDeleteFromArray(array [5]int, deleteVal int) [5]int {
	for i, val := range array {
		if val == deleteVal {
			for j := i; j < len(array)-1; j++ {
				array[j] = array[j+1]
			}
			break
		}
	}
	return array
}
func IndexOfLargestElementOfArray(array [5]int) int {
	if len(array) <= 0 {
		return -1
	}
	currLargeIdx := 0
	if len(array) == 1 {
		return 0
	}
	currLarge := array[0]

	for i := 1; i < len(array); i++ {
		if currLarge < array[i] {
			currLarge = array[i]
			currLargeIdx = i
		}
	}
	return currLargeIdx
}
func SecondLargestIndex(arr []int, n int) int {
	res, largest := -1, 0
	for i := 1; i < n; i++ {
		if arr[i] > arr[largest] {
			res = largest
			largest = i
		} else if arr[i] != arr[largest] {
			if res == -1 || arr[i] > arr[res] {
				res = i
			}
		}

	}
	return res
}

func IsArrayIsSorted(arr []int) bool {
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
func ReverseTheArray(arr []int) []int {
	low, high := 0, len(arr)-1
	for low < high {
		temp := arr[low]
		arr[low] = arr[high]
		arr[high] = temp
		high--
		low++
	}
	return arr
}

func RemoveDuplicatesFromSortedArray(arr []int, n int) []int {
	res := 1
	for i := 1; i < n; i++ {
		if arr[i] != arr[res-1] {
			arr[res] = arr[i]
			res++
		}
	}
	return arr
}

func MoveZerosToLast(arr []int) []int { //[]int{4, 0, 7, 0, 7, 7}
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
			if i != 0 {
				arr[i] = 0
			}
			count++
		}
	}
	return arr
}

func LeftRotateArrayByOne(arr []int) []int {
	temp := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = temp
	return arr
}

func PrintLeadersOfArray(arr []int) { //for present element there shouldnt be any greter element
	currLeader := arr[len(arr)-1] //[]int{1, 8, 3, 4, 5}
	fmt.Println(currLeader)
	for i := len(arr) - 2; i >= 0; i-- {
		if currLeader < arr[i] {
			currLeader = arr[i]
			fmt.Println(currLeader)
		}
	}
}

func MaximumDifferenceInArray(arr []int) int { //[]int{1, 8, 3, 6, 5}    o/p 7
	maxDiff, min := arr[1]-arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		maxDiff = int(math.Max(float64(arr[i]-min), float64(maxDiff)))
		min = int(math.Min(float64(min), float64(arr[i])))
	}
	return maxDiff
}

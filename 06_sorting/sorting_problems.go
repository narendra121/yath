package sorting

import (
	"fmt"
	"math"
	"sort"
)

// moving largest elements to the end
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		swapped := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			return arr
		}
	}
	return arr
}

// move smaller element to starting theta n^2
// not stable
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		min_idx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min_idx] {
				min_idx = j
			}
		}
		arr[min_idx], arr[i] = arr[i], arr[min_idx]
		fmt.Println(i, min_idx, arr)
	}
	return arr
}

// BigO(n)
// stable
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

// theta (m+n)
// sorting.MergeSort([]int{3, 3, 6, 9}, []int{1, 2, 3, 4, 5, 6}

func merge(left, right []int) []int {
	result := []int{}

	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func InterSection(arr1, arr2 []int) []int {
	l, r := 0, 0
	temp := make([]int, 0)
	for l < len(arr1) && r < len(arr2) {
		if arr1[l] > arr2[r] {

			r++
		} else if arr1[l] < arr2[r] {
			l++
		} else {
			temp = append(temp, arr1[l])
			l++
			r++
		}
	}
	return temp
}

func UnionOfArrays(arr1, arr2 []int) []int {
	l, r := 0, 0
	tmp := make([]int, 0)
	for (l < len(arr1)) && (r < len(arr2)) {
		if (l > 0) && (arr1[l] == arr1[l-1]) {
			l++
			continue
		}
		if r > 0 && (arr2[r] == arr2[r-1]) {
			r++
			continue
		}
		if arr1[l] < arr2[r] {
			tmp = append(tmp, arr1[l])
			l++
		} else if arr1[l] > arr2[r] {
			tmp = append(tmp, arr2[r])
			r++
		} else {
			tmp = append(tmp, arr1[l])
			l++
			r++
		}
	}
	for l < len(arr1) {
		if l > 0 && arr1[l] != arr1[l-1] {
			tmp = append(tmp, arr1[l])
		}
		if r > 0 && arr2[r] != arr2[r-1] {
			tmp = append(tmp, arr2[r])
		}
	}
	return tmp
}

// Lomuto Partition
// till i all elements less that pivot , after that it should >= piviot
func LomutoPartition(arr []int, low, height int) int {
	var piviot int
	fmt.Println(arr, low, height)

	if height == (len(arr) - 1) {
		piviot = arr[height]
	} else {
		arr[len(arr)-1], arr[height] = arr[height], arr[len(arr)-1]
	}
	j := low - 1

	for i := low; i <= height-1; i++ {
		if piviot > arr[i] {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j+1], arr[height] = arr[height], arr[j+1]
	return j + 1
}

// sorting.HoaresPartition([]int{5, 3, 8, 4, 2, 7, 1, 10}, 0, 7)
// 4
func HoaresPartition(arr []int, low, length int) int {
	piviot := arr[low]
	i, j := low-1, length+1
	for {
		for {
			i++
			if arr[i] > piviot {
				break
			}
		}
		for {
			j--
			if j < 0 || arr[j] < piviot {
				break
			}
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func QuickSortUsingLomuto(arr []int, l, h int) {
	if l < h {
		// fmt.Println("==", arr)
		// Choose a pivot element and partition the array around it.
		p := LomutoPartition(arr, l, h)
		// Recursively sort the left and right subarrays.
		QuickSortUsingLomuto(arr, l, p-1)
		QuickSortUsingLomuto(arr, p+1, h)
	}

}

// O(n2)
func KtSmallestElement(arr []int, n, k int) int {
	l, h := 0, n-1
	for l <= h {
		p := LomutoPartition(arr, l, h)
		if p == k-1 {
			return p
		} else if p > k-1 {
			h = p - 1
		} else {
			l = p + 1
		}
	}
	return -1
}

// inp  []int{2000, 49, 700, 500, 1}
// out 48
func MinimumDiffInArray(arr []int) int {
	sort.Ints(arr)
	res := math.MaxInt
	for i := 1; i < len(arr); i++ {
		res = int(math.Min(float64(res), float64(arr[i])-float64(arr[i-1])))
	}
	return res
}

/*
array is choco pkts and m is children number
pick given number of packts and there should minimum difference between min and max
in - []int{7, 3, 1, 8, 9, 12, 56}
out -2
O(nlogn)
*/
func ChocolateDistribution(arr []int, m int) int {
	if m > len(arr) {
		return -1
	}
	sort.Ints(arr)
	res := arr[m-1] - arr[0]
	for i := 1; (i + m - 1) < len(arr); i++ {
		res = int(math.Min(float64(res), float64(arr[i+m-1]-arr[i])))
	}
	return res
}

/*
in []int{-12, 18, -10, 15,-12}
out [-12 -12 -10 15 18]
uing HoaresPartition
teta(n)
*/
func SortPositioveAndNegetives(arr []int) []int {
	i, j := -1, len(arr)
	for {
		for {
			i++
			if arr[i] > 0 {
				break
			}
		}
		for {
			j--
			if j < 0 || arr[j] < 0 {
				break
			}
		}
		if i >= j {
			return arr
		}
		arr[i], arr[j] = arr[j], arr[i]
	}

}

type Interval struct {
	Start int
	End   int
}

func ComparatorSort(arr []Interval) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Start < arr[j].Start
	})
}

/*
O(nlogn)
input - []sorting.Interval{sorting.Interval{5, 10},sorting.Interval{3, 15},sorting.Interval{18, 30},sorting.Interval{2, 7}}
output  2 15

	18 30
*/
func MergeOverlapIntervals(arr []Interval) {
	ComparatorSort(arr)
	res := 0
	for i := 1; i < len(arr); i++ {
		if arr[res].End >= arr[i].Start {
			arr[res].End = int(math.Max(float64(arr[res].End), float64(arr[i].End)))
			arr[res].Start = int(math.Min(float64(arr[res].Start), float64(arr[i].Start)))
		} else {
			res++
			arr[res] = arr[i]
		}

	}
	for i := 0; i <= res; i++ {
		fmt.Println(arr[i].Start, arr[i].End)
	}

}

/*
in []int{900, 600, 700}, []int{1000, 800, 730}
out 2
*/
func MeetingMaxGuests(arrival, departure []int) int {
	sort.Ints(arrival)
	sort.Ints(departure)
	i, j, res, cur := 1, 0, 1, 1
	for i < len(arrival) && j < len(departure) {
		if arrival[i] <= departure[j] {
			cur++
			i++
		} else {
			cur--
			j++
		}
		res = int(math.Max(float64(res), float64(cur)))
	}
	return res
}

func BuildHeap(arr []int, len int) {
	for i := (len - 2) / 2; i >= 0; i-- {
		maxHeapify(arr, len, i)
	}
	// fmt.Println(arr)
}

func maxHeapify(arr []int, len, i int) {
	largest, left, rigt := i, 2*i+1, 2*i+2
	if left < len && arr[left] > arr[largest] {
		largest = left
	}
	if rigt <= len && arr[rigt] >= arr[largest] {
		largest = rigt
	}
	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		maxHeapify(arr, len, largest)
	}
}

func HeapSort(arr []int) {
	BuildHeap(arr, len(arr))
	fmt.Println(arr, len(arr))
	for i := len(arr) - 1; i >= 1; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapify(arr, i, 0)
	}
	fmt.Println(arr)
}

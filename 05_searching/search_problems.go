package searching

func BinararySearch(arr []int, low, high, searchItem int) int {
	for low <= high {
		mid := (high + low) / 2
		if arr[mid] == searchItem {
			return mid
		} else if arr[mid] > searchItem {
			high = mid - 1
		} else if arr[mid] < searchItem {
			low = mid + 1
		}
	}
	return -1
}

func RecursiveBinarySearch(arr []int, searchItem, low, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if arr[mid] == searchItem {
		return mid
	} else if arr[mid] > searchItem {
		return RecursiveBinarySearch(arr, searchItem, low, mid-1)
	} else if arr[mid] < searchItem {
		return RecursiveBinarySearch(arr, searchItem, mid+1, high)
	}
	return mid
}

func FirstOccuranceOfANumber(arr []int, searchItem int) int { // []int{5,10,10,20,20},20   0/p 3
	length := len(arr)
	low, high := 0, length-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] > searchItem {
			high = mid - 1
		} else if arr[mid] < searchItem {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != arr[mid] { //mid==0 because we are checking firstt occurance
				return mid
			} else {
				high = mid - 1
			}

		}
	}
	return -1
}

func LastOccuranceOccuranceOfANumber(arr []int, searchItem int) int {
	length := len(arr)
	low, high := 0, length-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] > searchItem {
			high = mid - 1
		} else if arr[mid] < searchItem {
			low = mid + 1
		} else {
			if mid == length-1 || arr[mid] != arr[mid+1] {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

func CountOfGivenDigit(arr []int, searchItem int) int {
	first := FirstOccuranceOfANumber(arr, searchItem)
	if first == -1 {
		return 0
	} else {
		return (LastOccuranceOccuranceOfANumber(arr, searchItem) - first + 1)
	}
}
func Count1sInSortedBinaryArray(arr []int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == 0 {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] == 0 {
				return ((len(arr)) - mid)
			}
			high = mid - 1
		}
	}
	return 0
}

func SearchSQRT(num int) int {
	low, high, ans := 1, num, -1
	for low < high {
		mid := (low + high) / 2
		tmp := mid * mid
		if tmp == num {
			return mid
		} else if tmp > num {
			high = mid - 1
		} else {
			low = mid + 1
			ans = mid
		}
	}
	return ans
}
func SearchInInfiniteArray(arr []int, searchItem int) int {
	if arr[0] == searchItem {
		return 0
	}
	i := 1
	for i <= len(arr)-1 && arr[i] < searchItem {
		if arr[i] == searchItem {
			return i
		}
		i = i * 2

	}
	return BinararySearch(arr, i/2+1, i-1, searchItem)
}

func SearchRotatedArray(arr []int, searchItem int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == searchItem {
			return mid
		} else if arr[low] <= arr[mid] {
			if arr[low] <= searchItem && arr[mid] >= searchItem {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if arr[mid] <= searchItem && arr[high] >= searchItem {
				low = mid + 1

			} else {
				high = mid - 1

			}
		}
	}
	return -1
}
func GetThePeakElement(arr []int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if (mid == 0 || arr[mid-1] <= arr[mid]) && (mid == len(arr)-1 || arr[mid+1] <= arr[mid]) {
			return mid
		}
		if arr[mid-1] >= arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

//two pointer approach
//any two elements in array are capable of making given sum
func IsAnySumEqual(arr []int, statingIdx, sum int) bool { //for set of Two element sum
	low, high := statingIdx, len(arr)-1
	for low < high {
		currSum := arr[low] + arr[high]
		if currSum == sum {
			return true
		} else if currSum > sum {
			high--
		} else {
			low++
		}
	}
	return false
}

func IsAnySumEqualTriplet(arr []int, sum int) bool {
	for i := 0; i < len(arr)-2; i++ {
		if IsAnySumEqual(arr, i+1, sum-arr[i]) {
			return true
		}
	}
	return false
}

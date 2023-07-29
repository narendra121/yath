package leetcode

import (
	"fmt"
	"sort"
)

func RomanToInt(s string) int {
	sl := make(map[byte]int)
	sl['I'] = 1
	sl['V'] = 5
	sl['X'] = 10
	sl['L'] = 50
	sl['C'] = 100
	sl['D'] = 500
	sl['M'] = 1000

	res := sl[s[0]]
	for i := 1; i <= len(s)-1; i++ {
		if sl[s[i]] > sl[s[i-1]] {
			res -= sl[s[i-1]]
			res += (sl[s[i]] - sl[s[i-1]])
		} else {
			res += sl[s[i]]

		}
	}
	return res

}
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		j := i + 1
		for j < len(nums) {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
			j++
		}
	}
	return []int{-1, -1}
}

func SingleNumber(nums []int) int { //([]int{2, 2, 2, 3}))
	sort.Ints(nums)
	if len(nums) == 1 {
		return nums[0]
	}
	fmt.Println(nums)
	if nums[len(nums)-1] != nums[len(nums)-2] {
		return nums[len(nums)-1]
	}
	if nums[0] != nums[1] {
		return nums[0]
	}
	for i := len(nums) - 2; i > 0; i-- {
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

/*

prev --- 2
     2 !=2 && 2!=2

prev 2
    2!=3 && 2 !=2

prev 3



*/

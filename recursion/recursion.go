package recursion

import (
	"fmt"
	"math"
	"strings"
)

func PrintNto1(num int) {
	if num <= 0 {
		return
	}
	fmt.Println(num)
	PrintNto1(num - 1)
}

func Print1toN(num int) {
	if num <= 0 {
		return
	}
	PrintNto1(num - 1)
	fmt.Println(num)
}

func FactorialOfNum(num, k int) int {
	if num <= 1 || k == 0 {
		return k
	}
	return FactorialOfNum(num-1, k*num)
}

func NthFibonacci(num int) int {
	if num <= 1 {
		return num
	}

	return NthFibonacci(num-1) + NthFibonacci(num-2)
}

func SumOfnNaturalNumbers(temp, num int) int {
	if num == 0 {
		return temp
	}
	return SumOfnNaturalNumbers(temp+num, num-1)
}

func PalindromeOrNot(inp string, start, end int) bool {
	if start >= end {
		return true
	}
	return (strings.Split(inp, "")[start] == strings.Split(inp, "")[end]) && PalindromeOrNot(inp, start+1, end-1)
}

func SumOfDigitsInNum(digiSum, inp int) int {
	if inp == 0 {
		return digiSum
	}
	return SumOfDigitsInNum(digiSum+(inp%10), inp/10)
}

func RopeCuttingProblem(ropeSize, size1, size2, size3 float64) float64 {
	if ropeSize == 0 {
		return 0
	} else if ropeSize < 0 {
		return -1
	}
	res := math.Max(math.Max(RopeCuttingProblem(ropeSize-size1, size1, size2, size3), RopeCuttingProblem(ropeSize-size2, size1, size2, size3)), RopeCuttingProblem(ropeSize-size2, size1, size2, size3))

	if res == -1 {
		return -1
	}
	return res + 1
}

func GeneratePossibleSubStrs(inp, res string, idx int) { //abc "" 0
	if idx == len(inp) {
		fmt.Println(res)
		return
	}
	GeneratePossibleSubStrs(inp, res, idx+1)
	GeneratePossibleSubStrs(inp, res+string(inp[idx]), idx+1) //abc
}

// https://www.youtube.com/watch?v=0nKIr3kAt-k
func TowerOfHanoi(height int, towerA, towerB, towerC string) {
	if height == 1 {
		fmt.Println("Move	1	from	" + towerA + "	to	" + towerC)
		return
	}
	TowerOfHanoi(height-1, towerA, towerC, towerB)
	fmt.Println("Move   ", height, "	from	"+towerA+"	to	"+towerC)
	TowerOfHanoi(height-1, towerB, towerA, towerC)

}

func JoshphusProblem(propleCount, position int) int {
	if propleCount == 1 {
		return 0
	}
	return (JoshphusProblem(propleCount-1, position) + position) % propleCount
}

func SubsetSumProblem(inp []int, length, sum int) int {
	if length == 0 {
		if sum == 0 {
			return 1
		}

		return 0

	}
	return SubsetSumProblem(inp, length-1, sum) + SubsetSumProblem(inp, length-1, sum-inp[length-1])
}

func Permutation(inp string, idx int) {
	if idx == len(inp)-1 {
		fmt.Println(inp)
		return
	}
	for i := idx; i < len(inp); i++ {
		temp := []byte(inp)
		temp[idx], temp[i] = temp[i], temp[idx]
		Permutation(string(temp), idx+1)
		temp1 := []byte(inp)
		temp1[idx], temp1[i] = temp1[i], temp1[idx]
	}
}

package bitmagic

import (
	"math"
	"sort"
	maths "yath/01_maths"
)

var bit32Table [256]int

func init() {
	bit32Table[0] = 0
	for i := 1; i < 256; i++ {
		bit32Table[i] = bit32Table[i&(i-1)] + 1
	}
}

func Factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * Factorial(num-1)
}

func GCD(num1, num2 int) int {
	if num2 == 0 {
		return num1
	} else {
		return GCD(num2, num1%num2)
	}
}

func LCM(num1, num2 int) int {
	ranger := int(math.Max(float64(num1), float64(num2)))
	for {
		if ranger%num1 == 0 && ranger%num2 == 0 {
			return ranger
		}
		ranger++
	}
}

func Divisors(num int) []int {
	result := make([]int, 0)
	result = append(result, 1)
	result = append(result, num)
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {

			result = append(result, i)
			if i != num/i {
				result = append(result, num/i)
			}
		}
	}
	sort.Ints(result)
	return result
}

func PrimeNumbers(num int) []int {
	result := make([]int, 0)
	for i := 1; i <= num; i++ {
		if maths.IsPrime(i) {
			result = append(result, i)
		}
	}
	return result
}

// 1 1 ->1  remaining 0
func BitWiseAnd(num1, num2 int) int {
	return num1 & num2
}

//0 0 --> 0  remaing 1
func BitWiseOr(num1, num2 int) int {
	return num1 | num2
}

//0 0 --> 0  1 1 ->0  remaing 1
func BitWiseXor(num1, num2 int) int {
	return num1 ^ num2
}

//num/2^shift
func BitWiseRightShift(num, shift int) int {
	return num >> shift
}

//num*2^shift
func BitWiseLeftShift(num, shift int) int {
	return num << shift
}

//Bitwise Not Available in golang
func BitWiseNotUint32(num int) uint32 {
	return ^uint32(num)
}

//Bitwise Not Available in golang
func BitWiseNotUint8(num int) uint8 {
	return ^uint8(num)
}

//number of 1's in binary
func SetBitsCount(num int) int {
	return bit32Table[num&255] + bit32Table[(num>>8)&255] + bit32Table[(num>>16)&255] + bit32Table[(num>>24)]
}

/*
	odd x^x^x^x=x
	even x^x^x=0
*/
func Repeat(sl []int, even bool) []int {
	res := make([]int, 0)
	checker := make(map[int]int)
	for _, val := range sl {
		checker[val] += 1
	}
	for key, val := range checker {
		if even && val%2 == 0 {
			res = append(res, key)
		} else if !even && val%2 != 0 {
			res = append(res, key)

		}
	}
	// res := sl[0]
	// for i := 1; i < len(sl); i++ {
	// 	res = res ^ sl[i]
	// }
	return res
}

package string_prblms

import (
	"fmt"
	"math"
	"strings"
	"yath/arrays"
	"yath/calc"
)

func CountEach(s string) {
	var charArr [26]int
	smallA := "a"
	s = strings.ToLower(s)
	for _, char := range s {
		charArr[int(char)-int('a')]++
	}

	for i := 0; i < 26; i++ {
		if charArr[i] > 0 {
			fmt.Println("count of "+string(int(smallA[0])+i), charArr[i])
		}
	}
}

func IsPalindrome(s string) bool {
	start, end := 0, len(s)-1

	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

//("geeksforgeeks", "grgesk")
//false
//("geeksforgeeks", "grgeks")
//true
func IsSubSequence(s, subSeqStr string) bool {
	strIdx, subSeqStrIdx := 0, 0
	for strIdx < len(s) && subSeqStrIdx < len(subSeqStr) {
		if s[strIdx] == subSeqStr[subSeqStrIdx] {
			strIdx++
			subSeqStrIdx++
		} else {
			strIdx++
		}
	}
	return subSeqStrIdx == len(subSeqStr)
}

//two strings should be having same characters, irrespective of order of characters

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	stringA := "a"
	var charArr [26]int
	for i := 0; i < len(s1); i++ {
		charArr[int(s1[i])-int(stringA[0])]++
		charArr[int(s2[i])-int(stringA[0])]--
	}
	for _, val := range charArr {
		if val != 0 {
			return false
		}
	}
	return true
}

/*
i/p:- geeksforgeeks
o/p:- 0

i/p:-abbcc
o/p: 1

i/p:- abcd
o/p -1
*/

func LeftMostRepeatingCharIdx(s string) int {

	charArray := make([]bool, 256)
	res := -1
	for i := len(s) - 1; i >= 0; i-- {
		visited := charArray[s[i]]
		if visited {
			res = i
		} else {
			charArray[s[i]] = true
		}
	}
	return res
}

func LeftMostNonRepeatingCharIdx(s string) int {
	charArray := arrays.FillSlice(256, -1)
	for i := 0; i < len(s); i++ {
		if charArray[s[i]] == -1 {
			charArray[s[i]] = i
		} else {
			charArray[s[i]] = -2
		}
	}
	// res := math.MaxInt
	for i := 0; i < len(s); i++ {
		if charArray[s[i]] >= 0 {
			return i
		}
	}
	return -1

}

//("abbcbda jam pom"
func ReverseWords(s string) string {
	start := 0

	for end := 0; end < len(s); end++ {
		if string(s[end]) == " " {
			Reverse(&s, start, end-1)
			start = end + 1
		}
	}
	Reverse(&s, start, len(s)-1)
	Reverse(&s, 0, len(s)-1)
	return s
}

func Reverse(s *string, low, high int) string {
	inp := []rune(*s)
	for low <= high {
		inp[low], inp[high] = inp[high], inp[low]
		low++
		high--
	}
	*s = string(inp)
	return *s
}

/*
INP :-ABABABCD
		ABAB
OUT:= 0,2

INP:- AAAAA
	AAA
OUT:- 0,1,2
*/
func PatternSearch(s, ptrn string) {
	lenOfTxt := len(s)
	lenOfPtrn := len(ptrn)
	for i := 0; i <= lenOfTxt-lenOfPtrn; i++ {
		var j int
		for j = 0; j < lenOfPtrn; j++ {
			if s[i+j] != ptrn[j] {
				break
			}
		}
		if j == lenOfPtrn {
			fmt.Println(i)
		}
	}
}

/*
INP:- ABCD
	  CDAB
	YES

*/
func IfStringsAreRotational(s1, s2 string) bool {
	return strings.Contains(s1+s1, s2)
}

/*
I/P:= geeksforgeeks
	  frog
O/P: true

I/P:= geeksforgeeks
	  rseek
O/P:=false
*/
func SearchAnagram(s, ptrn string) bool {
	var sArr [256]int
	var ptrnArr [256]int
	for i := 0; i < len(ptrn); i++ {
		sArr[s[i]]++
		ptrnArr[ptrn[i]]++
	}

	for i := len(ptrn); i < len(s); i++ {
		if CompareArrays(sArr, ptrnArr) {
			return true
		}
		sArr[s[i]]++
		sArr[s[i-len(ptrn)]]--
	}
	return false
}

func CompareArrays(x, y [256]int) bool {
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

/*
I/P:= BAC
O/P:= true

I/P:= CBA
O/P:=6
*/

func LexicoGraphicRank(s string) int {
	res := 1
	sLength := len(s)
	mul := calc.Factorial(sLength)
	count := make([]int, 256)
	for i := 0; i < sLength; i++ {
		count[s[i]]++
	}
	for i := 1; i < 256; i++ {
		count[i] += count[i-1]
	}

	for i := 0; i < sLength-1; i++ {
		mul = mul / (sLength - i)
		res = res + count[s[i]-1]*mul
		for j := int(s[i]); j < 256; j++ {
			count[j]--
		}
	}
	return res
}

/*
IP:= abcadbd
OP:= bcad

*/
func LongestSistinctSubStr(s string) int {
	sLength, res := len(s), 0
	prev := arrays.FillSlice(256, -1)
	i := 0
	for j := 0; j < sLength; j++ {
		i = int(math.Max(float64(i), float64(prev[s[j]]+1)))
		maxEnd := j - i + 1
		res = int(math.Max(float64(res), float64(maxEnd)))
		prev[s[j]] = j
	}
	return res
}

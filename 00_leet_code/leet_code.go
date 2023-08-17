package leetcode

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
	linkedList "yath/10_linkedList"
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

type HashTableSlice interface {
	AddValueToSlice(inp ...int)
	AddValueIndex(val int, index int)
	IsValueExist(val int) bool
	GetIndexOfValue(val int) int
	PrintAll()
}

type HashTable struct {
	mySlice []int
	myMap   map[int]int
}

func NewHashTable() HashTableSlice {
	return &HashTable{
		mySlice: make([]int, 0),
		myMap:   make(map[int]int),
	}
}

func (mt *HashTable) AddValueToSlice(inp ...int) {
	for _, val := range inp {
		temp := len(mt.mySlice)
		if mt.IsValueExist(val) {
			continue
		}
		mt.mySlice = append(mt.mySlice, val)
		if temp == 0 {
			temp = 0
		} else {
			temp -= 1
		}
		mt.AddValueIndex(val, temp)
	}
}

func (mt *HashTable) AddValueIndex(val int, index int) {
	mt.myMap[val] = index
}

func (mt *HashTable) IsValueExist(val int) bool {
	_, ok := mt.myMap[val]
	return ok
}

func (mt *HashTable) GetIndexOfValue(val int) int {
	idx, ok := mt.myMap[val]
	if !ok {
		return -1
	}
	return idx
}

func (mt *HashTable) PrintAll() {
	for _, val := range mt.mySlice {
		fmt.Println(val)
	}
}

var ladders = [][]int{{1, 4, 8, 21, 28, 50, 71, 80}, {38, 14, 30, 42, 76, 67, 92, 99}}
var snakes = [][]int{{32, 36, 48, 62, 88, 95, 97}, {10, 6, 26, 18, 24, 56, 78}}

func GetPosition(inp int, usrPos int) int {
	temp := usrPos
	usrPos += inp
	snake := 0
	ladder := 0
	if usrPos > 100 {
		fmt.Println("value exceeded max limit of 100")
		return temp
	}
	for j := 0; j < len(ladders[0]); j++ {
		if ladders[0][j] == usrPos {
			ladder = ladders[1][j]
		}
	}
	for j := 0; j < len(snakes[0]); j++ {
		if snakes[0][j] == usrPos {
			snake = snakes[1][j]
		}
	}
	if snake > 0 {
		return snake
	} else if ladder > 0 {
		return ladder
	} else {
		return usrPos
	}
}

func StartSnakesAndLaddersGame() {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator with the current time
	people := 0
	// generate a random integer between 1 and 6 (inclusive)
	fmt.Println("Enter Number of people: ")
	fmt.Scan(&people)
	dropperCount := people
	ScoreBoard := make([][]int, 2)
	for i := range ScoreBoard {
		ScoreBoard[i] = make([]int, people)
	}
	k := 0
	for j := 0; j < people; j++ {
		k += 1
		ScoreBoard[0][j] = k
	}

	fmt.Println(ScoreBoard)

	for {
		for j := 0; j < people; j++ {
			inp := rand.Intn(6) + 1
			check := ""
			if dropperCount == 1 {
				for j := 0; j < people; j++ {
					if ScoreBoard[1][j] >= 0 {
						fmt.Println()

						fmt.Println("THE WINNER IS ", ScoreBoard[0][j])
						return
					}
				}
			}
			for !strings.EqualFold(check, "y") || !strings.EqualFold(check, "d") || !strings.EqualFold(check, "s") {
				fmt.Println()
				if ScoreBoard[1][j] < 0 {
					fmt.Println("Player:", ScoreBoard[0][j], " Dropped from game")
					break
				}
				fmt.Println("Player:", ScoreBoard[0][j], " Please type :\n\tY: To roll the dice.\n\tS: To display scoreboard.\n\tD: To drop from the game.")
				fmt.Print("Type your choice: ")
				fmt.Scan(&check)
				if strings.EqualFold(check, "y") {
					fmt.Println("dice value is: ", inp)
					prevPosition := ScoreBoard[1][j]
					ScoreBoard[1][j] = GetPosition(inp, ScoreBoard[1][j])

					fmt.Println("position changed from: ", prevPosition, "to :", ScoreBoard[1][j])
					if ScoreBoard[1][j] == 100 {
						fmt.Println("THE WINNER IS Player: ", ScoreBoard[0][j])
						return
					}
					break
				} else if strings.EqualFold(check, "s") {
					fmt.Println(ScoreBoard)
				} else if strings.EqualFold(check, "d") {
					ScoreBoard[1][j] = -1
					fmt.Println("Dropping Player: ", ScoreBoard[0][j])
					dropperCount -= 1
				} else {
					fmt.Println("wrong input")

				}
			}
		}

	}
}

func AddTwoNumbers(l1 *linkedList.Node, l2 *linkedList.Node) *linkedList.Node {
	sum := getIntFromLinkedList(l1) + getIntFromLinkedList(l2)
	return CreateLinkedList(strconv.Itoa(sum), len(strconv.Itoa(sum))-1)
}

func getIntFromLinkedList(l *linkedList.Node) int {
	var reStr string
	if l == nil {
		return 0
	}
	reStr += strconv.Itoa(l.Data)
	temp := l
	for temp.Next != nil {
		reStr += strconv.Itoa(temp.Next.Data)
		temp = temp.Next
	}
	r, _ := strconv.Atoi(reStr)
	return r
}

func CreateLinkedList(sum string, idx int) *linkedList.Node {

	if idx < 0 {
		return nil
	}
	temp, _ := strconv.Atoi(string(sum[idx]))

	return &linkedList.Node{temp, CreateLinkedList(sum, idx-1)}
}

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



/*
 * Complete the 'find_max_elements' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY array as parameter.
 */

func find_max_elements(array []int32) int32 {
n:=array[0]
target:=array[1]

dp:=make([]int,target+1)
for  i:=int32(2);i<n-1;i++{
    num:=array[i]
    for i:=target;i>=num;i--{
        dp[i]=max(dp[i],dp[i-num]+1)
    }
    
}
return int32(dp[target])
}
func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    arrayCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var array []int32

    for i := 0; i < int(arrayCount); i++ {
        arrayItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        arrayItem := int32(arrayItemTemp)
        array = append(array, arrayItem)
    }

    result := find_max_elements(array)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

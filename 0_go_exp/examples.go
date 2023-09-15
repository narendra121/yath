package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
Identifiers
myVariable is a identifier
*/
var myIdentifier string

// myVariable it helps to store and manipulate data
var myVariable int = 5

// constants
const (
	Untype        = "hello"
	Typed  string = "hey"
)

// Conditional or decision
func IfExample(x int) {
	if x == 1 {
		fmt.Println("help")
	} else if x == 2 {
		fmt.Println("kill")
	} else {
		fmt.Println("scam")
	}
}

// loops
func LoopingExample(x []int) {
	//normal
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	//while with for
	i := 0
	for i == len(x)-1 {
		fmt.Println(x[i])
		i++
	}

	//infinite loop
	j := 0
	for {
		if j == len(x)-1 {
			break
		}
		fmt.Println(x[j])
		j++
	}

	//ranging
	for _, val := range x {
		fmt.Println(val)
	}

}

// switch example
func SwitchExample(x interface{}) {
	//switch case
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("default")
	}

	//type check
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("default")
	}
}

// variadic params
func VariadicParamExample(x ...int) {
	for _, val := range x {
		fmt.Println(val)
	}
}

// Anonymous func
func AnonymousFunction(myFunc func(x, y int) int) int {
	return myFunc(1, 2)
}

// defer example
func DeferExample() { //LIFO
	defer PrintMe("stop It")
	defer PrintMe("Narendra")
	defer PrintMe("Hello")

}

// struct
type GameInfo struct {
	GameName         string
	GameSquadCount   int
	GamePlayersCount int
	GameWonSquad     string
}

// interface
type Game interface {
	Name() string
	SquadCount() int
	PlayersCount() int
	Winner() string
}

// methods
func (p *GameInfo) Name() string {
	return p.GameName
}
func (p *GameInfo) SquadCount() int {
	return p.GameSquadCount
}

func (p *GameInfo) PlayersCount() int {
	return p.GamePlayersCount
}

func (p *GameInfo) Winner() string {
	return p.GameWonSquad
}

func Sender1() {
	fmt.Println("hello")
}
func SelectChecker(one, two, three, four, five chan int) {
	for i := 1; i <= 10; i++ {
		one <- i
		two <- i
		three <- i
		four <- i
		five <- i
	}
	close(one)
	close(two)
	close(three)
	close(four)
	close(five)
}
func main() {
	// go routine
	// go Sender1()
	// time.Sleep(3 * time.Second)

	// chan1 := make(chan int)
	// chan2 := make(chan int)
	// chan3 := make(chan int)
	// chan4 := make(chan int)
	// chan5 := make(chan int)
	// go SelectChecker(chan1, chan2, chan3, chan4, chan5)
	// for {
	// 	select {
	// 	case <-chan1:
	// 		fmt.Println("Im in channel 1")
	// 	case <-chan2:
	// 		fmt.Println("Im in channel 2")
	// 	case <-chan3:
	// 		fmt.Println("Im in channel 3")
	// 	case <-chan4:
	// 		fmt.Println("Im in channel 4")
	// 	case <-chan5:
	// 		fmt.Println("Im in channel 5")
	// 	default:
	// 		fmt.Println("Im in channel 6")

	// 	}
	// }
	// time.Sleep(5 * time.Minute)
	m := minMaxDifference(90)
	fmt.Println(m)
	// s:="qq"
	// strconv.Itoa(s)
	// a:=s[1]
	// strconv.Atoi(string(s[0]))
	// a := "soe"
	// for _, val = range strings.Split(a, "") {

	// }
	

}

func PrintMe(x interface{}) {
	fmt.Println(x)
}
func minMaxDifference(num int) int {
	var store [10]int
	tempNum := num
	for num > 0 {
		temp := num % 10
		num /= 10
		store[temp] += 1
	}
	maxIdx, maxVal, maxIdx1, maxVal1 := -1, -1, -1, -1

	for idx, val := range store {
		if val >= maxVal {
			maxVal = val
			maxIdx = idx
			maxIdx1 = idx
			maxVal1 = val
			fmt.Println(maxIdx, maxVal, maxIdx1, maxVal1)
		}

	}
	fmt.Println(maxIdx, maxVal)
	max, min := "", ""
	strNum := strconv.Itoa(tempNum)
	for i := 0; i < len(strNum); i++ {
		dig, _ := strconv.Atoi(string(strNum[i]))
		if dig == maxIdx {
			max += "9"
			min += "0"
		} else {
			max += string(strNum[i])
			min += string(strNum[i])
		}
	}
	small, _ := strconv.Atoi(min)
	big, _ := strconv.Atoi(max)

	if small == 0 {
		return big
	}
	return big - small
}

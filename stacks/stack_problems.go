package stacks

import "fmt"

/*
([]) -->true

((()) --> false

([)]  -->false

{}([()]) -->true

*/
func IsBalancedBrackets(str string) bool {
	var myStack Stack
	myStack.Init()
	for _, val := range str {
		if val == '{' || val == '(' || val == '[' {
			myStack.Push(string(val))
		} else {
			fmt.Println((myStack.Peek().(string) == string(val)))
			if myStack.IsEmpty() {
				return false
			} else if !(myStack.Peek().(string) == string(val)) {
				return false
			} else {
				myStack.Pop()
			}
		}
	}
	return myStack.IsEmpty()
}

func isMatching(s string) bool {
	return s == ")" || s == "]" || s == "}"
}

func StockSpan(arr []int, n int) {
	stock := StackSliceImpl
	stock.Push(0)
	for i := 1; i < n; i++ {
		for !stock.IsEmpty() && arr[stock.Peek().(int)] <= arr[i] {
			stock.Pop()
		}
		var span int
		if stock.IsEmpty() {
			span = i + 1
		} else {
			span = i - stock.Peek().(int)
		}
		fmt.Println(span)
		stock.Push(i)
	}
}

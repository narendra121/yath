package linkedList

import "fmt"

/*Node struct here
Data --> will store the value
Next --> will store Next node's address */
type Node struct {
	Data int
	Next *Node
}

//It will create the node from the input value
func GetNode(inp int) *Node {
	return &Node{
		Data: inp,
		Next: nil,
	}
}

//get simple linked list
func CreateSingleLinkedList() *Node {
	head := GetNode(10)
	head.Next = GetNode(21)
	head.Next.Next = GetNode(30)
	head.Next.Next.Next = GetNode(43)
	head.Next.Next.Next.Next = GetNode(50)
	head.Next.Next.Next.Next.Next = GetNode(60)

	return head
}

//Create a single linkedList from given array
func CreateSingleLinkedListFromArray(inp []int, idx int) *Node {
	if idx == len(inp) {
		return nil
	}
	return &Node{inp[idx], CreateSingleLinkedListFromArray(inp, idx+1)}
}

func TraverseSingleLinkedList(head *Node) {
	if head == nil {
		return
	}
	fmt.Println(head.Data)
	TraverseSingleLinkedList(head.Next)
}

func InsertAtTheBeginOfSingleLinkdList(head *Node, inp int) *Node { //O(n)
	temp := GetNode(inp)
	temp.Next = head
	return temp
}

func InsertAtTheEndOfLinkdList(head *Node, inp int) *Node { //O(n)
	temp := GetNode(inp)
	if head == nil {
		return temp
	}
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = temp
	return head
}
func DeleteHeadOfTheSingleLinkdList(head *Node) *Node {
	return head.Next
}

func DeleteTailOfTheSingleLinkdList(head *Node) *Node {
	cur := head
	if head == nil || head.Next == nil {
		return nil
	}
	for cur.Next.Next != nil {
		cur = cur.Next
	}
	cur.Next = nil
	return head
}
func InsertValueAtKthNodeSingleLinkdList(head *Node, position, Data int) *Node {
	temp := GetNode(Data)
	if position == 1 {
		temp.Next = head
		return temp
	}
	cur := head
	for i := 1; i < position && cur != nil; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return head
	}
	temp.Next = cur.Next
	return head
}

func SearchInSingleLinkdList(position int, head *Node) int {
	if position == 1 && head != nil {
		return head.Data
	}
	cur := head
	for i := 1; i < position && cur.Next != nil; i++ {
		cur = cur.Next
		if i < position && cur.Next == nil {
			return -1
		}
	}
	return cur.Data
}

//10 5 20 15
func CreateSingleCircularLinkdList() *Node {
	head := GetNode(10)
	temp1 := GetNode(5)
	temp2 := GetNode(20)
	temp3 := GetNode(15)
	head.Next = temp1
	head.Next.Next = temp2
	head.Next.Next.Next = temp3
	head.Next.Next.Next.Next = head
	return head
}

func TraverseSingleCircularLinkedList(head *Node) {

	if head == nil {
		return
	}
	fmt.Println(head.Data)
	for p := head.Next; p != head; p = p.Next {
		fmt.Println(p.Data)
	}
}

// https://github.com/onosproject/gnxi-simulators/blob/master/docs/README.md
func InsertAtBeginOfCircularSingleLinkedList(head *Node, Data int) *Node {

	temp := head
	newHead := GetNode(Data)
	newHead.Next = head
	if head == nil {
		temp.Next = temp
	} else {
		for temp.Next != head {
			temp = temp.Next
		}
		temp.Next = newHead
	}
	return newHead
}

func InsertAtTheEndOfTheCircularSingleLinkedList(head *Node, Data int) *Node {
	newTail := GetNode(Data)
	newTail.Next = head
	temp := head
	if head == nil {
		head.Next = newTail
	} else {
		for temp.Next != head {
			temp = temp.Next
		}
		temp.Next = newTail
	}
	return head
}
func DeleteKthNodeFromCircularLinkedList(head *Node, position int) *Node {
	temp := head
	if head == nil || position <= 0 {

	} else if position == 1 {
		for temp.Next != head {
			temp = temp.Next
		}
		temp.Next = head.Next
		head = head.Next
	} else {

		for i := 1; i < position-1 && temp != nil && temp != head; i++ {
			temp = temp.Next
		}
		oldTail := temp.Next
		temp.Next = oldTail.Next
		oldTail = nil
	}
	return head
}

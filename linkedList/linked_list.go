package linkedList

import "fmt"

/*Node struct here
data --> will store the value
next --> will store next node's address */
type Node struct {
	data int
	next *Node
}

//It will create the node from the input value
func getNode(inp int) *Node {
	return &Node{
		data: inp,
		next: nil,
	}
}

//get simple linked list
func CreateSingleLinkedList() *Node {
	head := getNode(10)
	head.next = getNode(20)
	head.next.next = getNode(30)
	head.next.next.next = getNode(40)
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
	fmt.Println(head.data)
	TraverseSingleLinkedList(head.next)
}

func InsertAtTheBeginOfSingleLinkdList(head *Node, inp int) *Node { //O(n)
	temp := getNode(inp)
	temp.next = head
	return temp
}

func InsertAtTheEndOfLinkdList(head *Node, inp int) *Node { //O(n)
	temp := getNode(inp)
	if head == nil {
		return temp
	}
	cur := head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = temp
	return head
}
func DeleteHeadOfTheSingleLinkdList(head *Node) *Node {
	return head.next
}

func DeleteTailOfTheSingleLinkdList(head *Node) *Node {
	cur := head
	if head == nil || head.next == nil {
		return nil
	}
	for cur.next.next != nil {
		cur = cur.next
	}
	cur.next = nil
	return head
}
func InsertValueAtKthNodeSingleLinkdList(head *Node, position, data int) *Node {
	temp := getNode(data)
	if position == 1 {
		temp.next = head
		return temp
	}
	cur := head
	for i := 1; i < position && cur != nil; i++ {
		cur = cur.next
	}
	if cur == nil {
		return head
	}
	temp.next = cur.next
	return head
}

func SearchInSingleLinkdList(position int, head *Node) int {
	if position == 1 && head != nil {
		return head.data
	}
	cur := head
	for i := 1; i < position && cur.next != nil; i++ {
		cur = cur.next
		if i < position && cur.next == nil {
			return -1
		}
	}
	return cur.data
}

//10 5 20 15
func CreateSingleCircularLinkdList() *Node {
	head := getNode(10)
	temp1 := getNode(5)
	temp2 := getNode(20)
	temp3 := getNode(15)
	head.next = temp1
	head.next.next = temp2
	head.next.next.next = temp3
	head.next.next.next.next = head
	return head
}

func TraverseSingleCircularLinkedList(head *Node) {

	if head == nil {
		return
	}
	fmt.Println(head.data)
	for p := head.next; p != head; p = p.next {
		fmt.Println(p.data)
	}
}

// https://github.com/onosproject/gnxi-simulators/blob/master/docs/README.md
func InsertAtBeginOfCircularSingleLinkedList(head *Node, data int) *Node {

	temp := head
	newHead := getNode(data)
	newHead.next = head
	if head == nil {
		temp.next = temp
	} else {
		for temp.next != head {
			temp = temp.next
		}
		temp.next = newHead
	}
	return newHead
}

func InsertAtTheEndOfTheCircularSingleLinkedList(head *Node, data int) *Node {
	newTail := getNode(data)
	newTail.next = head
	temp := head
	if head == nil {
		head.next = newTail
	} else {
		for temp.next != head {
			temp = temp.next
		}
		temp.next = newTail
	}
	return head
}
func DeleteKthNodeFromCircularLinkedList(head *Node, position int) *Node {
	temp := head
	if head == nil || position <= 0 {

	} else if position == 1 {
		for temp.next != head {
			temp = temp.next
		}
		temp.next = head.next
		head = head.next
	} else {

		for i := 1; i < position-1 && temp != nil && temp != head; i++ {
			temp = temp.next
		}
		oldTail := temp.next
		temp.next = oldTail.next
		oldTail = nil
	}
	return head
}

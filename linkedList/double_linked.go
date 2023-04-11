package linkedList

import "fmt"

type dNode struct {
	Data int
	Next *dNode
	prev *dNode
}

func getDNode(Data int) *dNode {
	return &dNode{
		Data: Data,
		Next: nil,
		prev: nil,
	}
}

func CreateDoubleLinkedList() *dNode {
	head := getDNode(10)
	temp1 := getDNode(20)
	temp2 := getDNode(30)
	head.Next = temp1
	temp1.prev = head
	temp1.Next = temp2
	temp2.prev = temp1
	return head
}

func CreateDoubleLinkedListFromArray(inp []int) {

}

func TraverseDoubleLinkedList(head *dNode) {
	if head == nil {
		fmt.Println("empty head")
		return
	}
	if head.Next == nil {
		fmt.Println(head.Data)
	}
	cur := head
	for cur.Next != nil {
		fmt.Println(cur.Data)
		cur = cur.Next
	}
	fmt.Println(cur.Data)
}

func InsertAtTheBeginOfDoubleLinkedList(head *dNode, Data int) *dNode {
	cur := getDNode(Data)
	cur.Next = head
	if head == nil {
		head.prev = cur
	}

	return cur
}

// 10 --> 20 -->  30 --> 40
// 40 --> 30 --> 20 10

func ReverseDoubleLinkedList(head *dNode) *dNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *dNode
	cur := head
	for cur != nil {
		prev = cur.prev
		cur.prev = cur.Next
		cur.Next = prev
		cur = cur.prev
	}
	return prev.prev
}

func DeleteTheHeadOfDoubleLinkdList(head *dNode) *dNode {
	if head == nil {
		return nil
	}
	newHead := head.Next
	if newHead != nil {
		newHead.prev = nil
	}
	head.Next = nil
	return newHead
}

func DeleteTailOfDoubleLinkdList(head *dNode) *dNode {
	if head == nil || head.Next == nil {
		return nil
	}
	// 10 20 30
	cur := head
	for cur.Next.Next != nil {
		cur = cur.Next
	}
	cur.Next = nil
	return head
}

func CreateCircularDoublyLinkedList() *dNode {
	head := getDNode(10)
	temp1 := getDNode(20)
	temp2 := getDNode(30)
	head.prev = temp2
	head.Next = temp1

	temp1.prev = head
	temp1.Next = temp2

	temp2.prev = temp1
	temp2.Next = head
	return head
}

func InsertAtTheBeginOfCircilarDoubleLinkedList(head *dNode, Data int) *dNode {
	newHead := getDNode(Data)
	headPrev := head.prev
	head.prev = newHead
	newHead.prev = headPrev
	newHead.Next = head
	headPrev.Next = newHead
	return newHead
}

func TraverseCircularDoublyLinkedList(head *dNode) {

	temp := head.Next

	fmt.Println(head.Data)
	for temp != head {
		fmt.Println(temp.Data)
		temp = temp.Next
	}

}

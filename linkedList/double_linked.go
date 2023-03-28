package linkedList

import "fmt"

type dNode struct {
	data int
	next *dNode
	prev *dNode
}

func getDNode(data int) *dNode {
	return &dNode{
		data: data,
		next: nil,
		prev: nil,
	}
}

func CreateDoubleLinkedList() *dNode {
	head := getDNode(10)
	temp1 := getDNode(20)
	temp2 := getDNode(30)
	head.next = temp1
	temp1.prev = head
	temp1.next = temp2
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
	if head.next == nil {
		fmt.Println(head.data)
	}
	cur := head
	for cur.next != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}
	fmt.Println(cur.data)
}

func InsertAtTheBeginOfDoubleLinkedList(head *dNode, data int) *dNode {
	cur := getDNode(data)
	cur.next = head
	if head == nil {
		head.prev = cur
	}

	return cur
}

// 10 --> 20 -->  30 --> 40
// 40 --> 30 --> 20 10

func ReverseDoubleLinkedList(head *dNode) *dNode {
	if head == nil || head.next == nil {
		return head
	}
	var prev *dNode
	cur := head
	for cur != nil {
		prev = cur.prev
		cur.prev = cur.next
		cur.next = prev
		cur = cur.prev
	}
	return prev.prev
}

func DeleteTheHeadOfDoubleLinkdList(head *dNode) *dNode {
	if head == nil {
		return nil
	}
	newHead := head.next
	if newHead != nil {
		newHead.prev = nil
	}
	head.next = nil
	return newHead
}

func DeleteTailOfDoubleLinkdList(head *dNode) *dNode {
	if head == nil || head.next == nil {
		return nil
	}
	// 10 20 30
	cur := head
	for cur.next.next != nil {
		cur = cur.next
	}
	cur.next = nil
	return head
}

func CreateCircularDoublyLinkedList() *dNode {
	head := getDNode(10)
	temp1 := getDNode(20)
	temp2 := getDNode(30)
	head.prev = temp2
	head.next = temp1

	temp1.prev = head
	temp1.next = temp2

	temp2.prev = temp1
	temp2.next = head
	return head
}

func InsertAtTheBeginOfCircilarDoubleLinkedList(head *dNode, data int) *dNode {
	newHead := getDNode(data)
	headPrev := head.prev
	head.prev = newHead
	newHead.prev = headPrev
	newHead.next = head
	headPrev.next = newHead
	return newHead
}

func TraverseCircularDoublyLinkedList(head *dNode) {

	temp := head.next

	fmt.Println(head.data)
	for temp != head {
		fmt.Println(temp.data)
		temp = temp.next
	}

}

package linkedList

import "fmt"

type DNode struct {
	Data int
	Next *DNode
	Prev *DNode
}

func GetDNode(Data int) *DNode {
	return &DNode{
		Data: Data,
		Next: nil,
		Prev: nil,
	}
}

func CreateDoubleLinkedList() *DNode {
	head := GetDNode(10)
	temp1 := GetDNode(20)
	temp2 := GetDNode(30)
	head.Next = temp1
	temp1.Prev = head
	temp1.Next = temp2
	temp2.Prev = temp1
	return head
}

func CreateDoubleLinkedListFromArray(inp []int) *DNode {
	prev := GetDNode(inp[0])
	head := prev
	for i, val := range inp {
		if i == 0 {
			continue
		}
		data := GetDNode(val)
		data.Prev = prev
		prev.Next = data
		prev = data
	}
	return head
}

func TraverseDoubleLinkedList(head *DNode) {
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

func InsertAtTheBeginOfDoubleLinkedList(head *DNode, Data int) *DNode {
	cur := GetDNode(Data)
	cur.Next = head
	if head == nil {
		head.Prev = cur
	}

	return cur
}

// 10 --> 20 -->  30 --> 40
// 40 --> 30 --> 20 10

func ReverseDoubleLinkedList(head *DNode) *DNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *DNode
	cur := head
	for cur != nil {
		prev = cur.Prev
		cur.Prev = cur.Next
		cur.Next = prev
		cur = cur.Prev
	}
	return prev.Prev
}

func DeleteTheHeadOfDoubleLinkdList(head *DNode) *DNode {
	if head == nil {
		return nil
	}
	newHead := head.Next
	if newHead != nil {
		newHead.Prev = nil
	}
	head.Next = nil
	return newHead
}

func DeleteTailOfDoubleLinkdList(head *DNode) *DNode {
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

func CreateCircularDoublyLinkedList() *DNode {
	head := GetDNode(10)
	temp1 := GetDNode(20)
	temp2 := GetDNode(30)
	head.Prev = temp2
	head.Next = temp1

	temp1.Prev = head
	temp1.Next = temp2

	temp2.Prev = temp1
	temp2.Next = head
	return head
}

func InsertAtTheBeginOfCircilarDoubleLinkedList(head *DNode, Data int) *DNode {
	newHead := GetDNode(Data)
	headPrev := head.Prev
	head.Prev = newHead
	newHead.Prev = headPrev
	newHead.Next = head
	headPrev.Next = newHead
	return newHead
}

func TraverseCircularDoublyLinkedList(head *DNode) {

	temp := head.Next

	fmt.Println(head.Data)
	for temp != head {
		fmt.Println(temp.Data)
		temp = temp.Next
	}

}

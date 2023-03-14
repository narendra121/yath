package linkedList

import "fmt"

type node struct {
	data int
	next *node
}

func getNode(inp int) *node {
	return &node{
		data: inp,
		next: nil,
	}
}

func LinkedList() *node {
	// var tempHead *node
	// for _, val := range inp {
	// 	if head == nil {

	// 	}
	// }
	head := getNode(10)
	head.next = getNode(20)
	head.next.next = getNode(30)
	head.next.next.next = getNode(40)
	// TraverseLinkedList(head)
	return head
}

func TraverseLinkedList(head *node) {
	temp := head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func RecursiveTraverseLinkedList(head *node) {
	if head == nil {
		return
	}
	fmt.Println(head.data)
	RecursiveTraverseLinkedList(head.next)
}

func InsertAtBeginLinkdList(head *node, inp int) *node { //O(n)
	temp := getNode(inp)
	temp.next = head
	return temp
}

func InsertAtEndOfLinkdList(head *node, inp int) *node { //O(n)
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

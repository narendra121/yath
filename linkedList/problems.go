package linkedList

func SortedInsertInLinkedList(head *Node, data int) {
	newNode := getNode(data)

	if head == nil {
		return head
	}
	if head.data > newNode.data {
		newNode.next = head
		return newNode
	}
	temp := head

	for temp.next != nil && temp.next.data < newNode.data {

		temp = temp.next

	}
	temp.next=newNode
}

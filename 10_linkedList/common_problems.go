package linkedList

func SortedInsertInLinkedList(head *Node, Data int) *Node {
	newNode := GetNode(Data)

	if head == nil {
		return head
	}
	if head.Data > newNode.Data {
		newNode.Next = head
		return newNode
	}
	temp := head

	for temp.Next != nil && temp.Next.Data < newNode.Data {

		temp = temp.Next

	}
	c := temp.Next.Next
	newNode.Next = c
	temp.Next = newNode
	return head
}

func GetMiddleOfLinkedList(head *Node) int {
	if head == nil {
		return 0
	}
	if head.Next == nil {
		return head.Data
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow.Data
}

func GetKthPositionFromEndInLinkedList(head *Node, position int) int {
	first, second := head, head
	if head == nil {
		return 0
	}
	for first.Next != nil && position > 0 {
		first = first.Next
		position -= 1
	}
	for first != nil {
		second = second.Next
		first = first.Next
	}
	return second.Data
	// for first.
}

func ReverseSingleLinkedList(head *Node) *Node {
	cur := head
	var prev *Node = nil
	for cur != nil {
		Next := cur.Next

		cur.Next = prev

		prev = cur
		cur = Next
	}
	return prev
}

func RemoveDuplicsFromSingleLinkdList(head *Node, present *Node) *Node {

	if present.Next == nil || head.Next == nil {
		return head
	}
	tempHead := head
	dupliCount := 0
	if tempHead.Data == present.Data {
		dupliCount++
	}
	for tempHead.Next != nil {
		if tempHead.Next.Data == present.Data {
			dupliCount++
		}
		if dupliCount > 1 {
			if tempHead.Next.Next == nil {
				break
			}
			newNext := tempHead.Next.Next
			tempHead.Next = newNext
		}

		tempHead = tempHead.Next
	}
	present = present.Next
	return RemoveDuplicsFromSingleLinkdList(head, present)
}

func ReverseSLinkedListAtKGroups(head *Node, k int) *Node {
	curr := head
	var prevFirst *Node
	isFirstHead := true
	for curr != nil {
		first := curr
		var prev *Node
		count := 0
		for curr != nil && count < k {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
			count++
		}
		if isFirstHead {
			head = prev
			isFirstHead = false
		} else {
			prevFirst.Next = prev
			prevFirst = first
		}
	}
	return head
}

func DetectLoopInSLinkdList(head *Node) bool {
	tortoise, flash := head, head
	for flash != nil && flash.Next != nil {
		tortoise = tortoise.Next
		flash = flash.Next.Next
		if tortoise == flash {
			return true
		}
	}
	return false
}

func DetectTheLoopAndReomve(head *Node) *Node {
	tortoise, flash := head, head
	for flash != nil && flash.Next != nil {
		tortoise = tortoise.Next
		flash = flash.Next.Next
		if tortoise == flash {
			break
		}
	}
	if flash != tortoise {
		return head
	}
	tortoise = head
	for tortoise.Next != flash.Next {
		tortoise = tortoise.Next
		flash = flash.Next

	}
	flash.Next = nil
	return head
}

func SegregateEvenAndOddInSLinkedList(head *Node) *Node {
	var evenStart, evenEnd, oddStart, oddEnd *Node
	for cur := head; cur != nil; cur = cur.Next {
		temp := cur.Data
		if temp%2 == 0 {
			if evenStart == nil {
				evenStart, evenEnd = cur, cur
			} else {
				evenEnd.Next = cur
				evenEnd = evenEnd.Next
			}
		} else {
			if oddStart == nil {
				oddStart, oddEnd = cur, cur
			} else {
				oddEnd.Next = cur
				oddEnd = oddEnd.Next
			}
		}
	}
	if oddStart == nil || evenStart == nil {

		return head
	}
	evenEnd.Next = oddStart
	oddEnd.Next = nil
	return evenStart
}

func LRUCacheDesign() { //least recently used
	
}

// 1 2 4

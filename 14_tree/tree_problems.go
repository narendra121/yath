package tree

import (
	"fmt"
	"math"
	stacks "yath/11_stacks"
	queue "yath/12_queue"
	dequeue "yath/13_deQueue"
)

type treeNode struct {
	key   interface{}
	left  *treeNode
	right *treeNode
}

func NewTreeNode(val interface{}) *treeNode {
	return &treeNode{
		key:   val,
		left:  &treeNode{},
		right: &treeNode{},
	}
}
func insert(arr []int, root *treeNode, i int, n int) *treeNode {
	if i < n {
		temp := &treeNode{key: arr[i], left: nil, right: nil}
		root = temp

		root.left = insert(arr, root.left, 2*i+1, n)
		root.right = insert(arr, root.right, 2*i+2, n)
	}
	return root
}

func CreateTree(arr []int) *treeNode {
	var root *treeNode
	root = insert(arr, root, 0, len(arr))
	return root
}

/*
InOrder (left,root,right)
PreOrder(Root,left,right)
post order(left,right,root)
*/

/*
 inp []int{1, 2, 3, 4, 5}
 out
4
2
5
1
3
*/
func InOrderTranversal(root *treeNode) {
	if root != nil {
		InOrderTranversal(root.left)
		fmt.Println(root.key)
		InOrderTranversal(root.right)
	}
}

/*
 inp []int{1, 2, 3, 4, 5}
 out
1
2
4
5
3
*/
func PreOrderTranversal(root *treeNode) {
	if root != nil {
		fmt.Println(root.key)
		PreOrderTranversal(root.left)
		PreOrderTranversal(root.right)
	}
}

/*
 inp []int{1, 2, 3, 4, 5}
 out
4
5
2
3
1

*/
func PostOrderTranversal(root *treeNode) {
	if root != nil {
		PostOrderTranversal(root.left)
		PostOrderTranversal(root.right)
		fmt.Println(root.key)
	}
}

/*
inp []int{1, 2, 3, 4, 5}
out  3

*/
func HeightOfTree(root *treeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(HeightOfTree(root.left)), float64(HeightOfTree(root.right)))) + 1
}

func BreadthFirstTraversal(root *treeNode) int {
	if root == nil {
		return 0
	}

	breadth := 0
	levelNodes := 1

	for levelNodes > 0 {
		nextLevelNodes := 0

		for i := 0; i < levelNodes; i++ {
			if root == nil {
				break
			}
			node := root

			if node.left != nil {
				nextLevelNodes++
			}
			if node.right != nil {
				nextLevelNodes++
			}
			root = root.right
		}

		breadth = int(math.Max(float64(breadth), float64(levelNodes)))
		levelNodes = nextLevelNodes
	}

	return breadth
}

/*
inp []int{1, 2, 3, 4, 5},2
out
		4 5
*/
func PrintElementsAtKthDistance(root *treeNode, k int) {
	if root == nil {
		return
	}
	if k == 0 {
		fmt.Println(root.key)
	} else {
		PrintElementsAtKthDistance(root.left, k-1)
		PrintElementsAtKthDistance(root.right, k-1)
	}
}

/*
inp:-[]int{20, 40, 30, 10, 60, 20, 40, 30, 10, 60}

out:-20      40      30      10      60      20      40      30      10      60
*/
func LevelOrderTraversal(root *treeNode) {
	q := queue.QueueArrImpl{}
	q.Init(0, 100)
	q.EnQueue(root)
	for q.Size() > 0 {
		curr := q.GetQueue().([]interface{})[0].(*treeNode)
		q.DeQueue()
		fmt.Print(curr.key, "\t")
		if curr.left != nil {
			q.EnQueue(curr.left)
		}
		if curr.right != nil {
			q.EnQueue(curr.right)
		}
	}
	fmt.Println()
}

/*
inp:=[]int{20, 40, 30, 10, 60, 20, 40, 30, 10, 60}
out:-
20
40      30
10      60      20      40
30      10      60

*/
func LevelOrderTraversalLineByLine(root *treeNode) {
	q := queue.QueueArrImpl{}
	q.Init(0, 100)
	q.EnQueue(root)
	q.EnQueue(nil)
	for q.Size() > 1 {
		var isNull bool
		switch q.GetQueue().([]interface{})[0].(type) {
		case nil:
			isNull = true
		}
		if isNull {
			fmt.Println()
			q.DeQueue()
			q.EnQueue(nil)
			continue
		}
		curr := q.GetQueue().([]interface{})[0].(*treeNode)
		q.DeQueue()

		fmt.Print(curr.key, "\t")
		if curr.left != nil {
			q.EnQueue(curr.left)
		}
		if curr.right != nil {
			q.EnQueue(curr.right)
		}
	}
	fmt.Println()
}

/*
inp:= ([]int{20, 40, 30, 10, 60, 20, 40, 30, 10, 60})
out:=10

*/
func SizeOfBinaryTree(root *treeNode) int {
	if root == nil {
		return 0
	}
	return 1 + SizeOfBinaryTree(root.left) + SizeOfBinaryTree(root.right)
}

/*
inp:=[]int{20, 40, 30, 10, 60, 20, 40, 30, 10, 60}
out :- 60
*/
func MaximumInBinaryTree(root *treeNode) int {
	if root == nil {
		return math.MinInt
	}
	return int(math.Max(float64(root.key.(int)), math.Max(float64(MaximumInBinaryTree(root.left)), float64(MaximumInBinaryTree(root.right)))))
}

/*
inp:= ([]int{10, 20, 30, 40, 50})
out:=
10
20
40
*/
func PrintLeftView(root *treeNode) {
	if root == nil {
		return
	}
	q := dequeue.DeQueueArrImpl{}
	q.Init(100)
	q.InsertFront(root)
	for !q.IsEmpty() {
		count := q.Size()
		for i := 0; i < count; i++ {
			curr := q.GetRear().(*treeNode)
			q.DeleteRear()

			if i == 0 {
				fmt.Println(curr.key)
			}
			if curr.left != nil {
				q.InsertFront(curr.left)
			}
			if curr.right != nil {
				q.InsertFront(curr.right)
			}
		}
	}
}

/*
inp:= []int{20, 12, 8, 3, 9}
out true

sum of  children values should be equals to root value
*/
func HasChildSumProperty(root *treeNode) bool {

	if root == nil {
		return true
	}
	if root.left == nil && root.right == nil {
		return true
	}
	sum := 0

	if root.left != nil {
		sum += root.left.key.(int)
	}
	if root.right != nil {
		sum += root.right.key.(int)
	}
	return (root.key == sum) && HasChildSumProperty(root.left) && HasChildSumProperty(root.right)
}

/*
the difference between height of left sub tree and right sub tree should not be more than 1

*/
func IsHeightBalanced(root *treeNode) int {
	if root == nil {
		return 0
	}
	lh := IsHeightBalanced(root.left)
	if lh == -1 {
		return -1
	}
	rh := IsHeightBalanced(root.right)
	if rh == -1 {
		return -1
	}
	if math.Abs(float64(lh)-float64(rh)) > 1 {
		return -1
	}
	return int(math.Max(float64(rh), float64(lh))) + 1
}

/*
inp:= []int{10, 20, 30, 40, 50, 60, 70}
out:= 4
*/
func Width(root *treeNode) int {
	if root == nil {
		return 0
	}
	dq := dequeue.DeQueueArrImpl{}
	dq.Init(100)
	dq.InsertFront(root)
	res := 0
	for !dq.IsEmpty() {
		count := dq.Size()
		res = int(math.Max(float64(res), float64(count)))
		for i := 0; i < count; i++ {
			curr := dq.GetRear().(*treeNode)
			dq.DeleteRear()
			if curr.left != nil {
				dq.InsertFront(curr.left)
			}
			if curr.right != nil {
				dq.InsertFront(curr.right)
			}
		}
	}
	return res
}

/*
inp:=[]int{10, 5, 20, 30, 50}
out= &{30 <nil> 0xc00000c048}

*/
var prev *treeNode

func TreeToDLL(root *treeNode) *treeNode {

	if root == nil {
		return root
	}
	head := TreeToDLL(root.left)
	if prev == nil {
		head = root
	} else {
		root.left = prev
		prev.right = root
	}
	prev = root
	TreeToDLL(root.right)
	return head
}

/*
we can construct tree if have below pair of arrays only
	inorder postorder
	inorder preorder

inp:= tree.ConstructTree([]int{20, 10, 40, 30, 50}, []int{10, 20, 30, 40, 50}, 0, 4)

*/
var preIdx int = 0

func ConstructTree(inOrder, preOrder []int, startIdx, endIdx int) *treeNode {
	if startIdx > endIdx {
		return nil
	}
	root := NewTreeNode(preOrder[preIdx])
	var inIdx int

	for i := startIdx; i <= endIdx; i++ {
		if inOrder[i] == root.key {
			inIdx = i
			break
		}
	}
	preIdx++
	root.left = ConstructTree(inOrder, preOrder, startIdx, inIdx-1)
	root.right = ConstructTree(inOrder, preOrder, inIdx+1, endIdx)
	return root
}

/*
inp:=[]int{20, 10, 40, 30, 50}
out:=
{20 0xc0000ba020 0xc0000ba080}
20
&{10 0xc0000ba040 0xc0000ba060} &{40 <nil> <nil>}
40
&{10 0xc0000ba040 0xc0000ba060}
10
&{50 <nil> <nil>} &{30 <nil> <nil>}
30
&{50 <nil> <nil>}
50
*/
func SpiralTraversal(root *treeNode) {
	stack1 := stacks.StackSliceImpl{}
	stack2 := stacks.StackSliceImpl{}
	stack1.Init()
	stack2.Init()
	stack1.Push(root)
	for !stack1.IsEmpty() || !stack2.IsEmpty() {
		for !stack1.IsEmpty() {
			tmp := stack1.Peek().(*treeNode)
			stack1.Pop()
			if tmp != nil {
				fmt.Println(tmp.key)
			}
			if tmp.left != nil {
				stack2.Push(tmp.left)
			}
			if tmp.right != nil {
				stack2.Push(tmp.right)
			}
		}
		for !stack2.IsEmpty() {
			tmp := stack2.Peek().(*treeNode)
			stack2.Pop()
			if tmp != nil {
				fmt.Println(tmp.key)
			}
			if tmp.right != nil {
				stack1.Push(tmp.right)
			}
			if tmp.left != nil {
				stack1.Push(tmp.left)
			}
		}
	}
}

/*
Number of nodes in maximum path in a tree between two leaf nodes

*/
var res int

func DiameterOfTree(root *treeNode) int {
	if root == nil {
		return 0
	}
	ld := DiameterOfTree(root.left)
	rd := DiameterOfTree(root.right)
	res = int(math.Max(float64(res), float64(1+ld+rd)))
	return int(1 + math.Max(float64(ld), float64(rd)))
}

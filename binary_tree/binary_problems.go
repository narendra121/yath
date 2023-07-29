package binarytree

import (
	"fmt"
	"math"
)

type binayTreeNode struct {
	key   int
	left  *binayTreeNode
	right *binayTreeNode
}

func insert(arr []int, root *binayTreeNode, i int, n int) *binayTreeNode {
	if i < n {
		temp := &binayTreeNode{key: arr[i], left: nil, right: nil}
		root = temp

		root.left = insert(arr, root.left, 2*i+1, n)
		root.right = insert(arr, root.right, 2*i+2, n)
	}
	return root
}

func CreateBinaryTree(arr []int) *binayTreeNode {
	var root *binayTreeNode
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
func InOrderTranversal(root *binayTreeNode) {
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
func PreOrderTranversal(root *binayTreeNode) {
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
func PostOrderTranversal(root *binayTreeNode) {
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
func HeightOfBinaryTree(root *binayTreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(HeightOfBinaryTree(root.left)), float64(HeightOfBinaryTree(root.right)))) + 1
}

/*
inp []int{1, 2, 3, 4, 5},2
out
		4 5
*/
func PrintElementsAtKthDistance(root *binayTreeNode, k int) {
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

func LevelOrderTraversal(root *binayTreeNode) {

}

package main

import (
	tree "yath/14_tree"
)

func main() {
	// dequeue.MaximumsInallSubArraysOfSizeK([]int{20, 40, 30, 10, 60}, 3)
	// root := tree.CreateTree([]int{10, 5, 20, 30, 50})
	// fmt.Println(tree.TreeToDLL(root))
	r := tree.CreateTree([]int{10, 20, 30, 40, 50})
	tree.IterativePreOrder(r)

	// tree.InOrderTranversal(tree.TreeToDLL(root))
}

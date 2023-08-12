package binarysearchtree

type BST struct {
	key   int
	left  *BST
	right *BST
}

func (b *BST) Init() {
	b.left = &BST{}
	b.right = &BST{}
}

func (b *BST) CreateNode(key int) {
	b.key = key
	b.left = &BST{}
	b.right = &BST{}
}

func (b *BST) Search(root *BST, x int) bool {
	if root == nil {
		return false
	} else if root.key == x {
		return true
	} else if root.key > x {
		return b.Search(root.left, x)
	} else {
		return b.Search(root.right, x)
	}
}

func (b *BST) IterativeSearch(root *BST, x int) bool {
	for root != nil {
		if root.key == x {
			return true
		} else if root.key < x {
			root = root.right
		} else {
			root = root.left
		}
	}
	return false
}

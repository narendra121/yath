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

func (b *BST) CreateNode(key int) *BST {
	return &BST{
		key:   key,
		left:  &BST{},
		right: &BST{},
	}
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

func (b *BST) Insert(root *BST, x int) *BST {
	if root == nil {
		return b.CreateNode(x)
	} else if root.key < x {
		root.right = b.Insert(root.right, x)
	} else if root.key > x {
		root.left = b.Insert(root.left, x)
	}
	return root
}

func (b *BST) IterativeInsert(root *BST, x int) *BST {
	temp := b.CreateNode(x)
	parent, curr := &BST{}, root
	for curr != nil {
		parent = curr
		if curr.key > x {
			curr = curr.left
		} else if curr.key < x {
			curr = curr.right
		} else {
			return root
		}
	}
	if parent == nil {
		return temp
	}
	if parent.key > x {
		parent.left = temp
	} else {
		parent.right = temp
	}
	return root
}

func DeleteNode(root *BST, x int) *BST {
	if root == nil {
		return root
	}
	if root.key > x {
		root.left = DeleteNode(root.left, x)

	} else if root.key < x {
		root.right = DeleteNode(root.right, x)
	} else {
		if root.left == nil {
			temp := root.right
			root = nil
			return temp
		} else if root.right == nil {
			temp := root.left
			root = nil
			return temp
		} else {
			succ := GetSuccessor(root)
			root.key = succ.key
			root.right = DeleteNode(root.right, succ.key)
		}
	}
	return root
}

func GetSuccessor(node *BST) *BST {
	node = node.right
	for node != nil && node.left != nil {
		node = node.left
	}
	return node
}

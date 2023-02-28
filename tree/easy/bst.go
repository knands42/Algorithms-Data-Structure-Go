package easy

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BST struct {
	root *Node
}

func (b *BST) Insert(val int) {
	if b.root == nil {
		b.root = &Node{Val: val}
		return
	}

	b.insert(val, b.root)
}

func (b *BST) insert(val int, node *Node) {
	if val < node.Val {
		if node.Left == nil {
			node.Left = &Node{Val: val}
		} else {
			b.insert(val, node.Left)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Val: val}
		} else {
			b.insert(val, node.Right)
		}
	}
}

func (b *BST) Search(val int) bool {
	return b.search(val, b.root)
}

func (b *BST) search(val int, node *Node) bool {
	if node == nil {
		return false
	}

	if val == node.Val {
		return true
	} else if val < node.Val {
		return b.search(val, node.Left)
	} else {
		return b.search(val, node.Right)
	}
}

func (b *BST) Delete(val int) {
	if b.root == nil {
		return
	}

	node := b.root
	var parent *Node

	for node != nil && node.Val != val {
		parent = node
		if val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		b.deleteLeafNode(node, parent)
	} else if node.Left != nil && node.Right != nil {
		b.deleteNodeWithTwoChildren(node, parent)
	} else {
		b.deleteNodeWithOneChild(node, parent)
	}
}

func (b *BST) deleteLeafNode(node *Node, parent *Node) {
	if parent == nil {
		b.root = nil
	} else if parent.Left == node {
		parent.Left = nil
	} else {
		parent.Right = nil
	}
}

func (b *BST) deleteNodeWithOneChild(node *Node, parent *Node) {
	if parent == nil {
		if node.Left != nil {
			b.root = node.Left
		} else {
			b.root = node.Right
		}
	} else if parent.Left == node {
		if node.Left != nil {
			parent.Left = node.Left
		} else {
			parent.Left = node.Right
		}
	} else {
		if node.Left != nil {
			parent.Right = node.Left
		} else {
			parent.Right = node.Right
		}
	}
}

func (b *BST) deleteNodeWithTwoChildren(node *Node, parent *Node) {
	predecessor, predecessorParent := b.getPredecessor(node.Left)
	temp := predecessor.Val
	predecessor.Val = node.Val
	node.Val = temp

	if predecessorParent == nil {
		if predecessor.Left != nil {
			node.Left = predecessor.Left
		} else {
			node.Left = nil
		}
		return
	}

	b.deleteLeafNode(predecessor, predecessorParent)
}

func (b *BST) getPredecessor(node *Node) (*Node, *Node) {
	var parent *Node
	for node.Right != nil {
		parent = node
		node = node.Right
	}

	return node, parent
}

func (b *BST) InOrder() []int {
	var stack []*Node
	var result []int
	node := b.root

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		node = node.Right
	}

	return result
}

func (b *BST) InOrderRecursion() []int {
	return b.inOrder(b.root)
}

func (b *BST) inOrder(node *Node) []int {
	if node == nil {
		return []int{}
	}

	return append(append(b.inOrder(node.Left), node.Val), b.inOrder(node.Right)...)
}

func (b *BST) Min() int {
	node := b.root
	for node.Left != nil {
		node = node.Left
	}

	return node.Val
}

func (b *BST) Max() int {
	node := b.root
	for node.Right != nil {
		node = node.Right
	}

	return node.Val
}

func (b *BST) Height() int {
	return b.height(b.root)
}

func (b *BST) height(node *Node) int {
	if node == nil {
		return 0
	}

	return 1 + max(b.height(node.Left), b.height(node.Right))
}

func (b *BST) Size() int {
	return b.size(b.root)
}

func (b *BST) size(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + b.size(root.Left) + b.size(root.Right)
}

func (b *BST) IsEmpty() bool {
	return b.root == nil
}

func (b *BST) IsBalanced() bool {
	return b.isBalanced(b.root)
}

func (b *BST) isBalanced(node *Node) bool {
	if node == nil {
		return true
	}

	leftHeight := b.height(node.Left)
	rightHeight := b.height(node.Right)

	return abs(leftHeight-rightHeight) <= 1 && b.isBalanced(node.Left) && b.isBalanced(node.Right)
}

func (b *BST) IsBST() bool {
	return b.isBST(b.root)
}

func (b *BST) isBST(node *Node) bool {
	if node == nil {
		return true
	}

	if node.Left != nil && node.Left.Val > node.Val {
		return false
	}

	if node.Right != nil && node.Right.Val < node.Val {
		return false
	}

	return b.isBST(node.Left) && b.isBST(node.Right)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

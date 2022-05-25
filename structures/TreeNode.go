package structures

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeNode(ar ...any) *TreeNode {
	if len(ar) == 0 {
		return nil
	}

	root := newNode(ar[0])
	queue := node{}
	queue.addLast(root)

	for i := 1; i < len(ar); {
		tn := queue.removeFirst()
		tn.Left = newNode(ar[i])
		i++
		if i < len(ar) {
			tn.Right = newNode(ar[i])
			i++
		}
		if tn.Left != nil {
			queue.addLast(tn.Left)
		}
		if tn.Right != nil {
			queue.addLast(tn.Right)
		}
	}
	return root
}

func newNode(x any) *TreeNode {
	v, ok := x.(int)
	if ok {
		return &TreeNode{Val: v}
	} else {
		return nil
	}
}

type node struct {
	val  *TreeNode
	next *node
}

func (n *node) addLast(tn *TreeNode) {
	lastnode := n
	for lastnode.next != nil {
		lastnode = lastnode.next
	}
	lastnode.next = &node{val: tn}
}

func (n *node) removeFirst() *TreeNode {
	if n.next == nil {
		return nil
	} else {
		tempN := n.next
		n.next = tempN.next
		return tempN.val
	}
}

package tree

import "sandbox/types"

type rbTNode[T types.Sortable] struct {
	parent      *rbTNode[T]
	children    []*rbTNode[T]
	val         T
	isBlackNode bool
}

type direction int

const (
	left = direction(iota)
	right
)

type RedBlackTree[T types.Sortable] struct {
	height int
	root   *rbTNode[T]
}

func (rb *RedBlackTree[T]) Insert(valuesToInsert ...T) {
	if len(valuesToInsert) == 0 {
		return
	}

	if rb.root == nil {
		rb.root = &rbTNode[T]{val: valuesToInsert[0], children: make([]*rbTNode[T], 2)}
		valuesToInsert = valuesToInsert[1:]
		rb.height++
	}

	for _, val := range valuesToInsert {
		rb.insert(rb.root, val)
	}
}

func (rb *RedBlackTree[T]) insert(node *rbTNode[T], val T) {
	nextDir := right
	if val < node.val {
		nextDir = left
	}

	if node.children[nextDir] == nil {
		node.children[nextDir] = &rbTNode[T]{val: val, parent: node, children: make([]*rbTNode[T], 2)}
		rb.balanceTreeFromNewLeaf(node.children[nextDir])
	} else {
		rb.insert(node.children[nextDir], val)
	}
}

func (rb *RedBlackTree[T]) balanceTreeFromNewLeaf(leaf *rbTNode[T]) {
	var grandParent, uncle *rbTNode[T]
	var dir direction
	node, parent := leaf, leaf.parent
	for parent != nil {
		if parent.isBlackNode {
			return
		}
		if parent.parent != nil {
			grandParent = parent.parent
			dir = parent.dirInParent()
			uncle = grandParent.children[1-dir]
		}

		if grandParent == nil {
			parent.isBlackNode = true
			return
		}

		if uncle == nil || uncle.isBlackNode {
			break
		}

		parent.isBlackNode = true
		uncle.isBlackNode = true
		grandParent.isBlackNode = false
		node = grandParent

		parent = node.parent // skip up one black level (two tree levels)
	}
	if parent == nil {
		return
	}

	if node == parent.children[1-dir] {
		parent.rotateDirRoot(rb, dir)
		node = parent
		parent = grandParent.children[dir]
	}
	grandParent.rotateDirRoot(rb, 1-dir)
	parent.isBlackNode = true
	grandParent.isBlackNode = false
}

func (node *rbTNode[T]) rotateDirRoot(rb *RedBlackTree[T], dir direction) {
	grandParent, oppositeDirChild := node.parent, node.children[1-dir]
	odcDirChild := oppositeDirChild.children[dir]

	node.children[1-dir] = odcDirChild
	if odcDirChild != nil {
		odcDirChild.parent = node
	}

	nodeDirInParent := node.dirInParent()
	oppositeDirChild.children[dir] = node
	node.parent = oppositeDirChild
	oppositeDirChild.parent = grandParent
	if grandParent != nil {
		grandParent.children[nodeDirInParent] = oppositeDirChild
	} else {
		rb.root = oppositeDirChild
	}
}

func (node *rbTNode[T]) dirInParent() direction {
	if node.parent.children[left] == node {
		return left
	}
	return right
}

package tree

import "sandbox/types"

type RedBlackTree[T types.Sortable] struct {
	height int
	root   *rbTNode[T]
}

func (rb *RedBlackTree[T]) Insert(valuesToInsert ...T) {
	if len(valuesToInsert) == 0 {
		return
	}

	if rb.root == nil {
		rb.root = newRbNode[T](valuesToInsert[0], nil)
		valuesToInsert = valuesToInsert[1:]
		rb.height++
	}

	for _, val := range valuesToInsert {
		rb.insert(rb.root, val)
	}
}

func (rb *RedBlackTree[T]) insert(node *rbTNode[T], val T) {
	nextDir := right
	if val < node.val() {
		nextDir = left
	}

	if node.childDir(nextDir) == nil {
		node.setChildDir(newRbNode[T](val, node), nextDir)
		rb.balanceTreeFromNewLeaf(node.castChildDir(nextDir))
	} else {
		rb.insert(node.castChildDir(nextDir), val)
	}
}

func (rb *RedBlackTree[T]) balanceTreeFromNewLeaf(leaf *rbTNode[T]) {
	var grandParent, uncle *rbTNode[T]
	var dir direction
	node, parent := leaf, leaf.castParent()
	for parent != nil {
		if parent.isBlackNode {
			return
		}
		if grandParent = parent.castParent(); grandParent != nil {
			dir = parent.dirInParent()
			uncle = grandParent.castChildDir(1 - dir)
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

		parent = node.castParent() // skip up one black level (two tree levels)
	}
	if parent == nil {
		return
	}

	if node == parent.castChildDir(1-dir) {
		parent.rotateDirRoot(rb, dir)
		node = parent
		parent = grandParent.castChildDir(dir)
	}
	grandParent.rotateDirRoot(rb, 1-dir)
	parent.isBlackNode = true
	grandParent.isBlackNode = false
}

func (node *rbTNode[T]) rotateDirRoot(rb *RedBlackTree[T], dir direction) {
	grandParent, oppositeDirChild := node.castParent(), node.castChildDir(1-dir)
	odcDirChild := oppositeDirChild.castChildDir(dir)

	node.setChildDir(odcDirChild, 1-dir)
	if odcDirChild != nil {
		odcDirChild.setParent(node)
	}

	var nodeDirInParent direction
	if grandParent != nil {
		nodeDirInParent = node.dirInParent()
	}
	oppositeDirChild.setChildDir(node, dir)
	node.setParent(oppositeDirChild)
	oppositeDirChild.setParent(grandParent)
	if grandParent != nil {
		grandParent.setChildDir(oppositeDirChild, nodeDirInParent)
	} else {
		rb.root = oppositeDirChild
	}
}

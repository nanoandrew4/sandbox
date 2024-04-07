package tree

import "sandbox/types"

type avlTNode[T types.Sortable] struct {
	parent, left, right *avlTNode[T]
	val                 T
	heightBelow         int
}

type AvlTree[T types.Sortable] struct {
	height int
	root   *avlTNode[T]
}

func (avl *AvlTree[T]) Insert(valuesToInsert ...T) {
	if len(valuesToInsert) == 0 {
		return
	}

	if avl.root == nil {
		avl.root = &avlTNode[T]{val: valuesToInsert[0]}
		valuesToInsert = valuesToInsert[1:]
		avl.height++
	}

	for _, val := range valuesToInsert {
		avl.height = avl.avlInsert(avl.root, val) + 1
	}
}

func (avl *AvlTree[T]) avlInsert(node *avlTNode[T], val T) int {
	if node == nil {
		return 0
	} else if val < node.val {
		if node.left == nil {
			node.left = &avlTNode[T]{val: val, parent: node}
			if node.right == nil {
				node.heightBelow++
			}
			return node.heightBelow
		}

		node.heightBelow = max(node.heightBelow, avl.avlInsert(node.left, val)+1)
	} else if node.val <= val {
		if node.right == nil {
			node.right = &avlTNode[T]{val: val, parent: node}
			if node.left == nil {
				node.heightBelow++
			}
			return node.heightBelow
		}

		node.heightBelow = max(node.heightBelow, avl.avlInsert(node.right, val)+1)
	}
	newSubtreeRoot := node.balanceIfNecessary()
	if newSubtreeRoot.parent == nil {
		avl.root = newSubtreeRoot
	}
	return node.heightBelow
}

func (an *avlTNode[T]) balanceIfNecessary() (newSubtreeRoot *avlTNode[T]) {
	bf := an.balanceFactor()
	if bf > 1 {
		if an.right.balanceFactor() > 0 { // right heavy
			return an.rotateLeft()
		} else {
			return an.rotateRightLeft()
		}
	}
	if bf < -1 { // left heavy
		if an.left.balanceFactor() < 0 {
			return an.rotateRight()
		} else {
			return an.rotateLeftRight()
		}
	}
	return an
}

func (an *avlTNode[T]) rotateLeft() (newSubtreeRoot *avlTNode[T]) {
	newSubtreeRoot = an.right
	if newSubtreeRoot.left != nil {
		newSubtreeRoot.left.parent = an
	}
	an.right = newSubtreeRoot.left
	newSubtreeRoot.left = an
	if an.parent != nil {
		swapChild(an.parent, an, newSubtreeRoot)
	} else {
		newSubtreeRoot.parent = nil
	}
	an.parent = newSubtreeRoot
	if newSubtreeRoot.balanceFactor() == 0 {
		newSubtreeRoot.heightBelow++
		an.heightBelow--
	} else {
		an.heightBelow -= 2
	}
	return newSubtreeRoot
}

func (an *avlTNode[T]) rotateRight() (newSubtreeRoot *avlTNode[T]) {
	newSubtreeRoot = an.left
	if newSubtreeRoot.right != nil {
		newSubtreeRoot.right.parent = an
	}
	an.left = newSubtreeRoot.right
	newSubtreeRoot.right = an
	if an.parent != nil {
		swapChild(an.parent, an, newSubtreeRoot)
	} else {
		newSubtreeRoot.parent = nil
	}
	an.parent = newSubtreeRoot
	if newSubtreeRoot.balanceFactor() == 0 {
		newSubtreeRoot.heightBelow++
		an.heightBelow--
	} else {
		an.heightBelow -= 2
	}
	return newSubtreeRoot
}

func (an *avlTNode[T]) rotateRightLeft() (newSubtreeRoot *avlTNode[T]) {
	oldRightChild := an.right
	newSubtreeRoot = oldRightChild.left
	newSubtreeRootRChild := newSubtreeRoot.right

	// right rotation
	oldRightChild.left = newSubtreeRootRChild
	if newSubtreeRootRChild != nil {
		newSubtreeRootRChild.parent = oldRightChild
	}
	newSubtreeRoot.right = oldRightChild
	newSubtreeRoot.right.parent = newSubtreeRoot

	// left rotation
	newSubtreeRootLChild := newSubtreeRoot.left
	an.right = newSubtreeRootLChild
	if newSubtreeRootLChild != nil {
		newSubtreeRootLChild.parent = an
	}
	newSubtreeRoot.left = an
	if an.parent != nil {
		swapChild(an.parent, an, newSubtreeRoot)
	} else {
		newSubtreeRoot.parent = nil
	}
	an.parent = newSubtreeRoot

	newSubtreeRoot.heightBelow++
	an.heightBelow--
	return newSubtreeRoot
}

func (an *avlTNode[T]) rotateLeftRight() (newSubtreeRoot *avlTNode[T]) {
	oldLeftChild := an.left
	newSubtreeRoot = oldLeftChild.right
	newSubtreeRootLChild := newSubtreeRoot.left

	// left rotation
	oldLeftChild.right = newSubtreeRootLChild
	if newSubtreeRootLChild != nil {
		newSubtreeRootLChild.parent = oldLeftChild
	}
	newSubtreeRoot.left = oldLeftChild
	newSubtreeRoot.left.parent = newSubtreeRoot

	// right rotation
	newSubtreeRootRChild := newSubtreeRoot.right
	an.left = newSubtreeRootRChild
	if newSubtreeRootRChild != nil {
		newSubtreeRootRChild.parent = an
	}
	newSubtreeRoot.right = an
	if an.parent != nil {
		swapChild(an.parent, an, newSubtreeRoot)
	} else {
		newSubtreeRoot.parent = nil
	}
	an.parent = newSubtreeRoot

	newSubtreeRoot.heightBelow++
	an.heightBelow--
	return newSubtreeRoot
}

func (an *avlTNode[T]) balanceFactor() int {
	var lh, rh int
	if an.left != nil {
		lh = an.left.heightBelow + 1
	}
	if an.right != nil {
		rh = an.right.heightBelow + 1
	}
	return rh - lh
}

func swapChild[T types.Sortable](parent, oldChild, newChild *avlTNode[T]) {
	if parent.left == oldChild {
		parent.left = newChild
	} else {
		parent.right = newChild
	}
}

func (avl *AvlTree[T]) Contains(val T) bool {
	_, node := findAvlParentAndNodeByVal(nil, avl.root, val)
	return node != nil
}

func findAvlParentAndNodeByVal[T types.Sortable](parent, node *avlTNode[T], val T) (p, n *avlTNode[T]) {
	if node == nil {
		return nil, nil
	} else if node.val == val {
		return parent, node
	} else if val < node.val {
		return findAvlParentAndNodeByVal(node, node.left, val)
	} else {
		return findAvlParentAndNodeByVal(node, node.right, val)
	}
}

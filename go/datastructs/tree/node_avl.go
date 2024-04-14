package tree

import "sandbox/types"

type avlTNode[T types.Sortable] struct {
	binaryTreeNodeWithParent[T]
	heightBelow int
}

func newAvlNode[T types.Sortable](val T, parent *avlTNode[T]) *avlTNode[T] {
	newNode := &avlTNode[T]{
		binaryTreeNodeWithParent: binaryTreeNodeWithParent[T]{
			binaryNode: &binaryTreeNode[T]{
				nodeVal: val,
			},
			nodeParent: parent,
		},
	}
	return newNode
}

func (node *avlTNode[T]) height() int {
	return node.heightBelow
}

func (node *avlTNode[T]) setHeight(newHeight int) {
	node.heightBelow = newHeight
}

func (node *avlTNode[T]) castLeft() *avlTNode[T] {
	return castAndReturnNode[T, *avlTNode[T]](node.left())
}

func (node *avlTNode[T]) castRight() *avlTNode[T] {
	return castAndReturnNode[T, *avlTNode[T]](node.right())
}

func (node *avlTNode[T]) castParent() *avlTNode[T] {
	return castAndReturnNode[T, *avlTNode[T]](node.parent())
}

func (node *avlTNode[T]) castChildDir(dir direction) *avlTNode[T] {
	return castAndReturnNode[T, *avlTNode[T]](node.childDir(dir))
}

func (node *avlTNode[T]) balanceIfNecessary() (newSubtreeRoot *avlTNode[T]) {
	bf := node.balanceFactor()
	if bf > 1 {
		if node.castRight().balanceFactor() > 0 { // right heavy
			return node.rotateDir(left)
		} else {
			return node.rotateOppositeDirAndThenDir(left)
		}
	}
	if bf < -1 { // left heavy
		if node.castLeft().balanceFactor() < 0 {
			return node.rotateDir(right)
		} else {
			return node.rotateOppositeDirAndThenDir(right)
		}
	}
	return node
}

func (node *avlTNode[T]) balanceFactor() int {
	var lh, rh int
	if lNode := node.castLeft(); lNode != nil {
		lh = lNode.height() + 1
	}
	if rNode := node.castRight(); rNode != nil {
		rh = rNode.height() + 1
	}
	return rh - lh
}

func (node *avlTNode[T]) rotateDir(dir direction) (newSubtreeRoot *avlTNode[T]) {
	newSubtreeRoot = node.castChildDir(1 - dir)
	newSubtreeRootChildDir := newSubtreeRoot.castChildDir(dir)
	if newSubtreeRootChildDir != nil {
		newSubtreeRootChildDir.setParent(node)
	}
	node.setChildDir(newSubtreeRootChildDir, 1-dir)
	newSubtreeRoot.setChildDir(node, dir)
	if node.castParent() != nil {
		swapChild(node.parent(), node, newSubtreeRoot)
	} else {
		newSubtreeRoot.setParent(nil)
	}
	node.setParent(newSubtreeRoot)
	if newSubtreeRoot.balanceFactor() == 0 {
		newSubtreeRoot.heightBelow++
		node.heightBelow--
	} else {
		node.heightBelow -= 2
	}
	return newSubtreeRoot
}

func (node *avlTNode[T]) rotateOppositeDirAndThenDir(dir direction) (newSubtreeRoot *avlTNode[T]) {
	oldRightChild := node.castChildDir(1 - dir)
	newSubtreeRoot = oldRightChild.castChildDir(dir)
	newSubtreeRootOppDirChild := newSubtreeRoot.castChildDir(1 - dir)

	// right rotation
	oldRightChild.setChildDir(newSubtreeRootOppDirChild, dir)
	if newSubtreeRootOppDirChild != nil {
		newSubtreeRootOppDirChild.setParent(oldRightChild)
	}
	newSubtreeRoot.setChildDir(oldRightChild, 1-dir)
	oldRightChild.setParent(newSubtreeRoot)

	// left rotation
	newSubtreeRootDirChild := newSubtreeRoot.castChildDir(dir)
	node.setChildDir(newSubtreeRootDirChild, 1-dir)
	if newSubtreeRootDirChild != nil {
		newSubtreeRootDirChild.setParent(node)
	}
	newSubtreeRoot.setChildDir(node, dir)
	if node.castParent() != nil {
		swapChild(node.parent(), node, newSubtreeRoot)
	} else {
		newSubtreeRoot.setParent(nil)
	}
	node.setParent(newSubtreeRoot)

	newSubtreeRoot.heightBelow++
	node.heightBelow--
	return newSubtreeRoot
}

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

func (node *avlTNode[T]) dirInParent() direction {
	if node.parent().left() == node {
		return left
	}
	return right
}

func (node *avlTNode[T]) balanceIfNecessary(avlTree *AvlTree[T]) {
	bf := node.balanceFactor()
	if bf > 1 {
		if node.castRight().balanceFactor() > 0 { // right heavy
			node.rotateDir(avlTree, left)
		} else {
			node.rotateOppositeDirAndThenDir(avlTree, left)
		}
	} else if bf < -1 { // left heavy
		if node.castLeft().balanceFactor() < 0 {
			node.rotateDir(avlTree, right)
		} else {
			node.rotateOppositeDirAndThenDir(avlTree, right)
		}
	}
}

func (node *avlTNode[T]) balanceFactor() int {
	var lh, rh int
	if lNode := node.castLeft(); lNode != nil {
		lh = lNode.heightBelow + 1
	}
	if rNode := node.castRight(); rNode != nil {
		rh = rNode.heightBelow + 1
	}
	return rh - lh
}

func (node *avlTNode[T]) rotateDir(avlTree *AvlTree[T], dir direction) *avlTNode[T] {
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
		avlTree.setRoot(oppositeDirChild)
	}

	node.heightBelow = getMaxHeightBelow(node.castLeft(), node.castRight()) + 1
	oppositeDirChild.heightBelow = getMaxHeightBelow(node, node.castChildDir(1-dir)) + 1
	return oppositeDirChild
}

func (node *avlTNode[T]) rotateOppositeDirAndThenDir(avlTree *AvlTree[T], dir direction) {
	node.castChildDir(1-dir).rotateDir(avlTree, 1-dir)
	node.rotateDir(avlTree, dir)
}

func getMaxHeightBelow[T types.Sortable](nodes ...*avlTNode[T]) int {
	maxHeight := -1 // if all nodes are nil, parent height will be calculated as 0
	for _, node := range nodes {
		if node != nil {
			maxHeight = max(maxHeight, node.heightBelow)
		}
	}
	return maxHeight
}

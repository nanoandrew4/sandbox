package tree

import "sandbox/types"

type rbTNode[T types.Sortable] struct {
	binaryTreeNodeWithParent[T]
	isBlackNode bool
}

func newRbNode[T types.Sortable](val T, parent *rbTNode[T]) *rbTNode[T] {
	newNode := &rbTNode[T]{
		binaryTreeNodeWithParent: binaryTreeNodeWithParent[T]{
			binaryNode: &binaryTreeNode[T]{
				nodeVal: val,
			},
			nodeParent: parent,
		},
	}
	return newNode
}

func (node *rbTNode[T]) dirInParent() direction {
	if node.parent().left() == node {
		return left
	}
	return right
}

func (node *rbTNode[T]) castParent() *rbTNode[T] {
	return castAndReturnNode[T, *rbTNode[T]](node.parent())
}

func (node *rbTNode[T]) castChildDir(dir direction) *rbTNode[T] {
	return castAndReturnNode[T, *rbTNode[T]](node.childDir(dir))
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
		rb.setRoot(oppositeDirChild)
	}
}

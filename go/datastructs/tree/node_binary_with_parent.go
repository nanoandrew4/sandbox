package tree

import "sandbox/types"

type binaryNodeWithParent[T types.Sortable] interface {
	binaryNode[T]
	parent() binaryNodeWithParent[T]
	setParent(binaryNodeWithParent[T])
	dirInParent() direction
}

type binaryTreeNodeWithParent[T types.Sortable] struct {
	binaryNode[T]
	nodeParent binaryNodeWithParent[T]
}

func (node *binaryTreeNodeWithParent[T]) parent() binaryNodeWithParent[T] {
	return node.nodeParent
}

func (node *binaryTreeNodeWithParent[T]) setParent(parent binaryNodeWithParent[T]) {
	node.nodeParent = parent
}

func (node *binaryTreeNodeWithParent[T]) dirInParent() direction {
	if node.parent().left() == node {
		return left
	}
	return right
}

func rotateDir[T types.Sortable, R binaryNodeWithParent[T]](node R, tree BinaryTree[T, R], dir direction) R {
	grandParent, oppositeDirChild := node.parent(), castAndReturnNode[T, R](node.childDir(1-dir))
	odcDirChild := oppositeDirChild.childDir(dir)

	node.setChildDir(odcDirChild, 1-dir)
	if !isNodeNil(odcDirChild) {
		odcDirChild.(R).setParent(node)
	}

	var nodeDirInParent direction
	if !isNodeNil[T](grandParent) {
		nodeDirInParent = node.dirInParent()
	}
	oppositeDirChild.setChildDir(node, dir)
	node.setParent(oppositeDirChild)
	oppositeDirChild.setParent(grandParent)
	if !isNodeNil[T](grandParent) {
		grandParent.setChildDir(oppositeDirChild, nodeDirInParent)
	} else {
		tree.setRoot(oppositeDirChild)
	}
	return oppositeDirChild
}

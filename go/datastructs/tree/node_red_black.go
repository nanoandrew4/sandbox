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
	return castAndReturnRbNode[T](node.parent())
}

func (node *rbTNode[T]) castChildDir(dir direction) *rbTNode[T] {
	return castAndReturnRbNode(node.childDir(dir))
}

func castAndReturnRbNode[T types.Sortable](node binaryNode[T]) *rbTNode[T] {
	if isNodeNil[T](node) {
		return nil
	}
	return node.(*rbTNode[T])
}

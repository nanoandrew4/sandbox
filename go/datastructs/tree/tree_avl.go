package tree

import "sandbox/types"

type AvlTree[T types.Sortable] struct {
	baseSortedBinaryTree[T, *avlTNode[T]]
	height int
}

func (avl *AvlTree[T]) Insert(valuesToInsert ...T) {
	if len(valuesToInsert) == 0 {
		return
	}

	if avl.root() == nil {
		avl.setRoot(newAvlNode[T](valuesToInsert[0], nil))
		valuesToInsert = valuesToInsert[1:]
		avl.height++
	}

	for _, val := range valuesToInsert {
		avl.height = avl.avlInsert(avl.root(), val) + 1
	}
}

func (avl *AvlTree[T]) avlInsert(node *avlTNode[T], val T) int {
	if val < node.val() {
		if node.castLeft() == nil {
			node.setLeft(newAvlNode[T](val, node))
			if node.castRight() == nil {
				node.setHeight(node.height() + 1)
			}
			return node.height()
		}

		node.setHeight(max(node.height(), avl.avlInsert(node.castLeft(), val)+1))
	} else if node.val() <= val {
		if node.castRight() == nil {
			node.setRight(newAvlNode[T](val, node))
			if node.castLeft() == nil {
				node.setHeight(node.height() + 1)
			}
			return node.height()
		}

		node.setHeight(max(node.height(), avl.avlInsert(node.castRight(), val)+1))
	}
	newSubtreeRoot := node.balanceIfNecessary()
	if newSubtreeRoot.castParent() == nil {
		avl.setRoot(newSubtreeRoot)
	}
	return node.height()
}

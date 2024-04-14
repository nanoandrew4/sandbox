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
		avl.avlInsert(avl.root(), val)
		avl.height = avl.root().heightBelow + 1
	}
}

func (avl *AvlTree[T]) avlInsert(node *avlTNode[T], val T) int {
	if val < node.val() {
		if node.castLeft() == nil {
			node.setLeft(newAvlNode[T](val, node))
			if node.castRight() == nil {
				node.heightBelow++
			}
			return node.heightBelow
		}

		node.heightBelow = max(node.heightBelow, avl.avlInsert(node.castLeft(), val)+1)
	} else if node.val() <= val {
		if node.castRight() == nil {
			node.setRight(newAvlNode[T](val, node))
			if node.castLeft() == nil {
				node.heightBelow++
			}
			return node.heightBelow
		}

		node.heightBelow = max(node.heightBelow, avl.avlInsert(node.castRight(), val)+1)
	}
	node.balanceIfNecessary(avl)
	return node.heightBelow
}

package datastructs

type TreeType int

const (
	preOrderTree = TreeType(iota)
	inOrderTree
	postOrderTree
)

type NodeTraversalFunc[T comparable] func(node *btNode[T]) (continueTraversal bool)

type btNode[T comparable] struct {
	left, right *btNode[T]
	val         T
}

type BinaryTree[T comparable] struct {
	tType TreeType
	root  *btNode[T]
}

func (bt *BinaryTree[T]) Insert(val T) {
	if bt.root == nil {
		bt.root = &btNode[T]{val: val}
		return
	}

	bt.walkBreadthFirstAndRunFunc(&Queue[*btNode[T]]{}, bt.root, func(node *btNode[T]) bool {
		if node.left == nil {
			node.left = &btNode[T]{val: val}
			return false
		} else if node.right == nil {
			node.right = &btNode[T]{val: val}
			return false
		}
		return true
	})
}

func (bt *BinaryTree[T]) WalkDepthFirst(f NodeTraversalFunc[T]) {
	bt.walkDepthFirstAndRunFunc(bt.root, f)
}

func (bt *BinaryTree[T]) walkDepthFirstAndRunFunc(node *btNode[T], f NodeTraversalFunc[T]) {
	if node == nil {
		return
	}

	switch bt.tType {
	case preOrderTree:
		f(node)
		bt.walkDepthFirstAndRunFunc(node.left, f)
		bt.walkDepthFirstAndRunFunc(node.right, f)
	case inOrderTree:
		bt.walkDepthFirstAndRunFunc(node.left, f)
		f(node)
		bt.walkDepthFirstAndRunFunc(node.right, f)
	case postOrderTree:
		bt.walkDepthFirstAndRunFunc(node.left, f)
		bt.walkDepthFirstAndRunFunc(node.right, f)
		f(node)
	}
}

func (bt *BinaryTree[T]) walkBreadthFirstAndRunFunc(q *Queue[*btNode[T]], node *btNode[T], f NodeTraversalFunc[T]) {
	if node == nil {
		return
	}

	q.Enqueue(node.left)
	q.Enqueue(node.right)
	if !f(node) {
		return
	}
	nextNode, _ := q.Dequeue() // if dequeue fails, nextNode should be nil, which will be caught with the first if statement
	bt.walkBreadthFirstAndRunFunc(q, nextNode, f)
}

func (bt *BinaryTree[T]) Equals(bt2 *BinaryTree[T]) bool {
	if bt2 == nil {
		return false
	}
	if !areNodesEqual(bt.root, bt2.root) {
		return false
	}
	return true
}

func areNodesEqual[T comparable](n1, n2 *btNode[T]) bool {
	if n1 == nil && n2 == nil {
		return true
	} else if n1 == nil || n2 == nil {
		return false
	} else if n1.val != n2.val {
		return false
	}
	return areNodesEqual(n1.left, n2.left) && areNodesEqual(n1.right, n2.right)
}

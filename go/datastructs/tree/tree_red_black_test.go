package tree

import (
	"math/rand"
	"sandbox/types"
	"testing"
)

const (
	insertionsToTest = 100000
)

func TestRedBlackTree_Insert(t *testing.T) {
	rbTree := &RedBlackTree[int]{}
	rbTree.Insert(50, 25, 75, 12, 37, 6, 18, 3)

	if rbTree.root().val() != 25 {
		t.Fatal("expected tree root to be 25")
	}
	if rbTree.root().castChildDir(left).val() != 12 {
		t.Fatal("expected tree root left child to be 12")
	}

	rbTree = &RedBlackTree[int]{}
	for range insertionsToTest {
		valToInsert := rand.Int()
		rbTree.Insert(valToInsert)
	}

	rbOrderedArray := rbTree.ToOrderedArray()
	for i := 1; i < len(rbOrderedArray); i++ {
		if rbOrderedArray[i-1] > rbOrderedArray[i] {
			t.Fatalf("rb tree ordered array check failed at index %d", i-1)
		}
	}

	checkRbTreeColoring(t, rbTree)
}

func checkRbTreeColoring[T types.Sortable](t *testing.T, tree *RedBlackTree[T]) {
	TraverseBreadthFirst(tree.root(), func(node binaryNode[T]) (continueTraversal bool) {
		castedNode := node.(*rbTNode[T])
		lChild, rChild := castedNode.castChildDir(left), castedNode.castChildDir(right)
		if !castedNode.isBlackNode && ((lChild != nil && !lChild.isBlackNode) || (rChild != nil && !rChild.isBlackNode)) {
			t.Fatal("red node has red children")
		}

		if (lChild != nil && rChild == nil && lChild.isBlackNode) || (lChild == nil && rChild != nil && rChild.isBlackNode) {
			t.Fatal("node has single black child")
		}

		return true
	})
}

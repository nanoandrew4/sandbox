package tree

import (
	"math/rand"
	rand2 "math/rand/v2"
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

func TestRedBlackTree_Delete(t *testing.T) {
	rbTree := &RedBlackTree[int]{}
	for range insertionsToTest {
		valToInsert := rand.Int()
		rbTree.Insert(valToInsert)
	}

	rbArray := rbTree.ToOrderedArray()
	rand2.Shuffle(len(rbArray), func(i, j int) {
		rbArray[i], rbArray[j] = rbArray[j], rbArray[i]
	})

	for idx, treeVal := range rbArray {
		if !rbTree.Contains(treeVal) {
			t.Fatalf("%d: expected tree to contain %d", idx, treeVal)
		}
		if !rbTree.Delete(treeVal) {
			t.Fatalf("%d: expected to delete %d successfully", idx, treeVal)
		}
	}
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

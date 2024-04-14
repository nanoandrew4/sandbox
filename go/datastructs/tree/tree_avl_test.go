package tree

import (
	"math/rand/v2"
	"testing"
)

func TestAvlTree_SingleRotationInsert(t *testing.T) {
	avlTree := &AvlTree[int]{}
	avlTree.Insert(1, 2, 3, 4, 5, 0, -1)

	if avlTree.height != 3 {
		t.Fatal("expected tree height of 3")
	}
	if avlTree.root().balanceFactor() > 1 || avlTree.root().balanceFactor() < -1 {
		t.Fatal("expected tree to be balanced")
	}
}

func TestAvlTree_DoubleRotationInsert(t *testing.T) {
	avlTree := &AvlTree[int]{}
	avlTree.Insert(5, 15, 1, 10, 20, 11)

	if avlTree.height != 3 {
		t.Fatal("expected tree height of 3")
	}
	if avlTree.root().val() != 10 {
		t.Fatal("expected tree root after double left rotation to be 10")
	}
	if avlTree.root().balanceFactor() > 1 || avlTree.root().balanceFactor() < -1 {
		t.Fatal("expected tree to be balanced")
	}

	avlTree = &AvlTree[int]{}
	avlTree.Insert(15, 20, 5, 3, 12, 13)
	if avlTree.height != 3 {
		t.Fatal("expected tree height of 3")
	}
	if avlTree.root().val() != 12 {
		t.Fatal("expected tree root after double right rotation to be 12")
	}
	if avlTree.root().balanceFactor() > 1 || avlTree.root().balanceFactor() < -1 {
		t.Fatal("expected tree to be balanced")
	}
}

func TestAvlTree_Insert(t *testing.T) {
	avlTree := &AvlTree[int]{}
	for range insertionsToTest {
		valToInsert := rand.Int()
		avlTree.Insert(valToInsert)
	}

	avlOrderedArray := avlTree.ToOrderedArray()
	for i := 1; i < len(avlOrderedArray); i++ {
		if avlOrderedArray[i-1] > avlOrderedArray[i] {
			t.Fatalf("avl tree ordered array check failed at index %d", i-1)
		}
	}

	TraversePreOrder(avlTree.root(), func(node binaryNode[int]) (continueTraversal bool) {
		castedNode := node.(*avlTNode[int])
		if castedNode.castParent() != nil && castedNode.castParent().heightBelow-castedNode.heightBelow > 2 {
			t.Fatal("height difference between nodes was greater than 2")
		}
		return true
	})
}

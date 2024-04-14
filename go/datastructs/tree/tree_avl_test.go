package tree

import "testing"

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

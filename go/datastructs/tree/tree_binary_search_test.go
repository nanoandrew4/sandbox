package tree

import (
	"slices"
	"testing"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	bst := &BinarySearchTree[int]{}
	bst.Insert(1, 2, 2, 3, 0)
	bst.Insert(-5)
	bst.Insert(-2)

	testArraysEqual(t, bst.ToOrderedArray(), []int{-5, -2, 0, 1, 2, 2, 3})
}

func TestBinarySearchTree_Contains(t *testing.T) {
	bst := &BinarySearchTree[int]{}
	bst.Insert(1)
	bst.Insert(0)
	bst.Insert(-2)

	if !bst.Contains(-2) {
		t.Fatal("expected binary search tree to contain -2")
	}
	if bst.Contains(-1) {
		t.Fatal("expected binary search tree to not contain -1")
	}
}

func testArraysEqual[T comparable](t *testing.T, actual, expected []T) {
	if !slices.Equal(actual, expected) {
		t.Errorf("expected %+v to equal %+v", actual, expected)
	}
}

func TestBinarySearchTree_Delete(t *testing.T) {
	bst := &BinarySearchTree[int]{}
	bst.Insert(1)
	if !bst.Delete(1) || bst.Contains(1) || bst.root() != nil {
		t.Fatal("expected 1 root node to be deleted")
	}

	bst.Insert(50, 25, 25, 75, 12, 37, 62, 87, 6, 18, 31, 43, 55, 68, 81, 93, 33, 32, 34, 5, 4, 3)
	// check preconditions
	if bst.root().left().val() != 25 {
		t.Fatal("expected left value from root to be 25")
	}
	if bst.root().left().right().val() != 25 {
		t.Fatal("expected replacement value for first 25 when its deleted to be second 25")
	}
	if bst.root().left().right().right().left().val() != 31 {
		t.Fatal("expected replacement value for second 25 when its deleted to be 31")
	}
	if bst.root().right().right().val() != 87 {
		t.Fatal("expected right value right child of root to be 87")
	}
	if bst.root().right().right().left().val() != 81 {
		t.Fatal("expected rightmost left leaf to be 81")
	}
	if bst.root().right().right().right().val() != 93 {
		t.Fatal("expected rightmost leaf to be 93")
	}

	if !bst.Delete(25) || !bst.Contains(25) || bst.root().left().val() != 25 || bst.root().left().right().left().val() != 31 {
		t.Fatal("expected first 25 to be deleted from tree and replaced by second 25")
	}
	if !bst.Delete(25) || bst.Contains(25) || bst.root().left().val() != 31 || bst.root().left().right().left().val() != 33 {
		t.Fatal("expected second 25 to be deleted from tree and replaced by 31")
	}
	if !bst.Delete(6) || bst.Contains(6) || bst.root().left().left().left().val() != 5 || bst.root().left().left().left().left().val() != 4 || bst.root().left().left().left().right() != nil {
		t.Fatal("expected 6 to be deleted from tree and replaced by 5, which has a left child 4 and no right child")
	}
	if !bst.Delete(81) || bst.Contains(81) || bst.root().right().right().left() != nil {
		t.Fatal("expected leaf 81 to be deleted from tree")
	}
	if !bst.Delete(93) || bst.Contains(93) || bst.root().right().right().right() != nil {
		t.Fatal("expected leaf 93 to be deleted from tree")
	}

	if bst.Delete(500) {
		t.Fatal("expected 500 not to be deleted from tree as it does not exist")
	}
}

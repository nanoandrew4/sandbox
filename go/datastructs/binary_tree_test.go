package datastructs

import (
	"testing"
)

func TestBinaryTree_Insert(t *testing.T) {
	bt := &BinaryTree[int]{}
	bt.Insert(1, 2, 3)

	var preOrderValues []int
	bt.WalkDepthFirst(func(node *btNode[int]) (continueTraversal bool) {
		preOrderValues = append(preOrderValues, node.val)
		return true
	})

	testArraysEqual(t, preOrderValues, []int{1, 2, 3})

	bt.tType = inOrderTree
	bt.Insert(4)

	var inOrderValues []int
	bt.WalkDepthFirst(func(node *btNode[int]) (continueTraversal bool) {
		inOrderValues = append(inOrderValues, node.val)
		return true
	})

	testArraysEqual(t, inOrderValues, []int{4, 2, 1, 3})

	bt.tType = postOrderTree

	var postOrderValues []int
	bt.WalkDepthFirst(func(node *btNode[int]) (continueTraversal bool) {
		postOrderValues = append(postOrderValues, node.val)
		return true
	})

	testArraysEqual(t, postOrderValues, []int{4, 2, 3, 1})
}

func TestBinaryTree_Equals(t *testing.T) {
	bt1 := &BinaryTree[int]{}
	bt1.Insert(1)
	bt1.Insert(2)
	bt1.Insert(3)

	bt2 := &BinaryTree[int]{}
	bt2.Insert(1)
	bt2.Insert(2)

	if bt1.Equals(nil) {
		t.Fatal("binary trees should not be equal when tree to compare against is nil")
	}

	if bt1.Equals(bt2) {
		t.Fatal("binary trees should not be equal")
	}

	bt2.Insert(3)
	if !bt1.Equals(bt2) {
		t.Fatal("binary trees should be equal")
	}

	// Same shape as 1, different values
	bt3 := &BinaryTree[int]{}
	bt3.Insert(1)
	bt3.Insert(2)
	bt3.Insert(4)

	if bt1.Equals(bt3) {
		t.Fatal("binary trees 1, 3 should not be equal")
	}
}

func TestBinaryTree_Contains(t *testing.T) {
	bt := &BinaryTree[int]{}
	bt.Insert(1)
	bt.Insert(2)
	bt.Insert(3)

	if !bt.Contains(2) {
		t.Fatal("expected tree to contain value 2")
	}

	if bt.Contains(5) {
		t.Fatal("expected tree to not contain value 5")
	}
}

func TestBinaryTree_Delete(t *testing.T) {
	bt := &BinaryTree[int]{}
	bt.Insert(1)
	if !bt.Delete(1) || bt.Contains(1) || bt.root != nil {
		t.Fatal("expected 1 root node to be deleted")
	}

	for i := 0; i < 100; i++ {
		bt.Insert(i)
	}

	// check preconditions
	if bt.root.left.val != 1 {
		t.Fatal("expected left value from root to be 1")
	}
	if bt.root.right.right.val != 6 {
		t.Fatal("expected right value right child of root to be 6")
	}

	if !bt.Delete(1) || bt.Contains(1) || bt.root.left.val == 1 {
		t.Fatal("expected 1 to be deleted from tree")
	}

	bt.root.right.right.left = nil // to force right side to bubble up
	if !bt.Delete(6) || bt.Contains(6) || bt.root.left.val == 6 {
		t.Fatal("expected 6 to be deleted from tree")
	}

	if !bt.Delete(61) || bt.Contains(61) {
		t.Fatal("expected 61 leaf node to be deleted, and parent left node set to nil")
	}
	if !bt.Delete(62) || bt.Contains(62) {
		t.Fatal("expected 61 leaf node to be deleted, and parent right node set to nil")
	}

	if bt.Delete(500) {
		t.Fatal("expected 500 not to be deleted from tree as it does not exist")
	}
}

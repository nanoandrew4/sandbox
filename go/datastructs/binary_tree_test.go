package datastructs

import (
	"slices"
	"testing"
)

func TestBinaryTree_Insert(t *testing.T) {
	bt := &BinaryTree[int]{}
	bt.Insert(1)
	bt.Insert(2)
	bt.Insert(3)

	var preOrderValues []int
	bt.WalkDepthFirst(func(node *btNode[int]) (continueTraversal bool) {
		preOrderValues = append(preOrderValues, node.val)
		return true
	})

	if !slices.Equal(preOrderValues, []int{1, 2, 3}) {
		t.Error("expected pre order values to be 1, 2, 3")
	}

	bt.tType = inOrderTree
	bt.Insert(4)

	var inOrderValues []int
	bt.WalkDepthFirst(func(node *btNode[int]) (continueTraversal bool) {
		inOrderValues = append(inOrderValues, node.val)
		return true
	})

	if !slices.Equal(inOrderValues, []int{4, 2, 1, 3}) {
		t.Error("expected in order values to be 4, 2, 1, 3")
	}

	bt.tType = postOrderTree

	var postOrderValues []int
	bt.WalkDepthFirst(func(node *btNode[int]) (continueTraversal bool) {
		postOrderValues = append(postOrderValues, node.val)
		return true
	})

	if !slices.Equal(postOrderValues, []int{4, 2, 3, 1}) {
		t.Error("expected in order values to be 4, 2, 3, 1")
	}
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
		t.Error("binary trees should not be equal when tree to compare against is nil")
	}

	if bt1.Equals(bt2) {
		t.Error("binary trees should not be equal")
	}

	bt2.Insert(3)
	if !bt1.Equals(bt2) {
		t.Error("binary trees should be equal")
	}

	// Same shape as 1, different values
	bt3 := &BinaryTree[int]{}
	bt3.Insert(1)
	bt3.Insert(2)
	bt3.Insert(4)

	if bt1.Equals(bt3) {
		t.Error("binary trees 1, 3 should not be equal")
	}
}

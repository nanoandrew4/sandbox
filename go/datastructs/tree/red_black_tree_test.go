package tree

import "testing"

func TestRedBlackTree_Insert(t *testing.T) {
	rbTree := &RedBlackTree[int]{}
	rbTree.Insert(50, 25, 75, 12, 37, 6, 18, 3)

	if rbTree.root.val != 50 {
		t.Fatal("expected tree root to be 50")
	}
	if rbTree.root.children[left].val != 12 {
		t.Fatal("expected tree root left child to be 12")
	}
}

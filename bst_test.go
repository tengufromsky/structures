package main

import (
	"testing"
)

func TestBST_InOrderDo(t *testing.T) {
	var bst BST
	src := []int{1, 10, 30, 100, 200, 300}
	for _, k := range src {
		bst.Insert(k, nil)
	}

	var rangedElems []int
	bst.InOrderDo(func(n *Node) {
		rangedElems = append(rangedElems, n.key)
	})
	if !isIntSlicesEqual(src, rangedElems) {
		t.Error("slices don't equal")
	}
}

func isIntSlicesEqual(a, b []int) bool {
	if len(a) < len(b) || len(a) > len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if i >= len(b)-1 {
			break
		}
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

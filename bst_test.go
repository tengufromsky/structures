package main

import (
	"fmt"
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

func TestBST_Delete(t *testing.T) {
	t.Run("balanced tree", func(t *testing.T) {
		var bst BST
		src := []int{100, 10, 200, 150, 125, 165, 1, 5, 2}
		for _, k := range src {
			bst.Insert(k, nil)
		}

		bst.Delete(2)
		bst.Delete(10)
		bst.Delete(200)
		bst.Delete(5)
		bst.Delete(150)

		var rangedElems []int
		bst.InOrderDo(func(n *Node) {
			rangedElems = append(rangedElems, n.key)
		})

		fmt.Println(rangedElems)
		if !isIntSlicesEqual([]int{1, 100, 125, 165}, rangedElems) {
			t.Error("slices don't equal")
		}
	})
	t.Run("disbalanced tree", func(t *testing.T) {
		var bst BST
		src := []int{1, 2, 5, 10, 100, 125, 150, 165, 200}
		for _, k := range src {
			bst.Insert(k, nil)
		}

		bst.Delete(2)
		bst.Delete(10)
		bst.Delete(200)
		bst.Delete(5)
		bst.Delete(150)

		var rangedElems []int
		bst.InOrderDo(func(n *Node) {
			rangedElems = append(rangedElems, n.key)
		})

		if !isIntSlicesEqual([]int{1, 100, 125, 165}, rangedElems) {
			t.Error("slices don't equal")
		}
	})
}

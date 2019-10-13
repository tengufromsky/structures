package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
	value interface{}
}

func (n *Node) isLeaf() bool {
	if n.left == nil && n.right == nil {
		return true
	}
	return false
}

type BST struct {
	root *Node
}

func (t *BST) Search(key int) *Node {
	if t.root == nil || t.root.key == key {
		return t.root
	}

	return t.searchNode(t.root, key)
}

func (t *BST) searchNode(node *Node, key int) *Node {
	if node == nil || node.key == key {
		return node
	}

	if key < node.key {
		return t.searchNode(node.left, key)
	}

	return t.searchNode(node.right, key)
}

func (t *BST) Insert(key int, value interface{}) {
	if t.root == nil {
		t.root = &Node{
			key:   key,
			value: value,
		}
		return
	}

	t.insertNode(t.root, key, value)
}

func (t *BST) insertNode(node *Node, key int, value interface{}) {
	if key == node.key {
		node.value = value
	}

	if key < node.key {
		if node.left == nil {
			node.left = &Node{
				key:   key,
				value: value,
			}
		} else {
			t.insertNode(node.left, key, value)
		}
		return
	}

	if node.right == nil {
		node.right = &Node{
			key:   key,
			value: value,
		}
	} else {
		t.insertNode(node.right, key, value)
	}
}

func (t *BST) InOrderDo(f func(n *Node)) {
	if t.root != nil {
		t.inOrderDo(t.root, f)
	}
}

func (t *BST) inOrderDo(node *Node, f func(n *Node)) {
	if node != nil {
		t.inOrderDo(node.left, f)
		f(node)
		t.inOrderDo(node.right, f)
	}
}

func (t *BST) Delete(key int) {
	t.root = t.deleteNode(t.root, key)
}

func (t *BST) deleteNode(root *Node, key int) *Node {
	if root == nil {
		return root
	}

	switch {
	case key < root.key:
		root.left = t.deleteNode(root.left, key)
	case key > root.key:
		root.right = t.deleteNode(root.right, key)
	default:
		// root is leaf
		if root.left == nil && root.right == nil {
			return nil // current node will be NILed in parent
		}

		// root have 1 child
		if root.left == nil {
			return root.right
		}
		if root.right == nil {
			return root.left
		}

		// root have two child
		if root.key == 150 {
			fmt.Println("dilemma here")
		}
		minValue := t.lessThanNode(root)
		root.key, root.value = minValue.key, minValue.value
		root.left = t.deleteNode(root.left, root.key) // replace new root from child
	}
	return root
}

// minValueNode return the node  with minum key value found in that tree.
func (t *BST) lessThanNode(node *Node) *Node {
	current := node
	if current == nil {
		return nil
	}

	for current.left != nil {
		current = current.left
	}

	fmt.Println(current)
	return current
}

package main

type Node struct {
	key   int
	left  *Node
	right *Node
	value interface{}
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

func (t *BST) inOrderDo(node *Node, f func(n *Node)) {
	if node != nil {
		t.inOrderDo(node.left, f)
		f(node)
		t.inOrderDo(node.right, f)
	}
}

func (t *BST) InOrderDo(f func(n *Node)) {
	if t.root != nil {
		t.inOrderDo(t.root, f)
	}
}

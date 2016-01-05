package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

func new_node(val int) *node {
	n := new(node)
	n.data = val
	n.left = nil
	n.right = nil
	return n
}

func (n *node) print() {
	fmt.Println(n.data)
}

type bst struct {
	root *node
}

func (tree *bst) insert(val int) {
	temp := tree.root
	if temp == nil {
		n := new_node(val)
		tree.root = n
		return
	}
	parent := temp
	for temp != nil {
		if temp.data > val {
			parent = temp
			temp = temp.left
		} else {
			parent = temp
			temp = temp.right
		}
	}
	n := new_node(val)
	if parent.data > val {
		parent.left = n
	} else {
		parent.right = n
	}
}

func (tree *bst) inorder() {
	inorder_helper(tree.root)
}
func inorder_helper(n *node) {
	if n == nil {
		return
	}
	inorder_helper(n.left)
	fmt.Println(n.data)
	inorder_helper(n.right)
}

func (tree *bst) height() int {
	return height_helper(tree.root)
}

func height_helper(n *node) int {
	if n == nil {
		return 0
	}
	left := height_helper(n.left)
	right := height_helper(n.right)
	if left > right {
		return 1 + left
	}
	return 1 + right
}

func main() {
	tree := bst{root: nil}
	tree.insert(5)
	tree.insert(3)
	tree.insert(1)
	tree.insert(7)
	tree.insert(8)
	tree.insert(6)
	tree.inorder()
	fmt.Println("height of the tree is ", tree.height())
}

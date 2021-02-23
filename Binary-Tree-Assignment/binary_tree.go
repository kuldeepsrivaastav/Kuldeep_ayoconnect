package Binary_Tree_Assignment

import (
	"fmt"
)

type Tree struct {
	Root *Node
}

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (t *Tree) Insert(data int) {
	if t.Root == nil {
		t.Root = &Node{key: data}
	} else {
		t.Root.insert(data)
	}
}

func (n *Node) insert(data int) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func PrintPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.key)
		PrintPreOrder(n.left)
		PrintPreOrder(n.right)
	}
}

func PrintPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		PrintPostOrder(n.left)
		PrintPostOrder(n.right)
		fmt.Printf("%d ", n.key)
	}
}

func PrintInOrder(n *Node) {
	if n == nil {
		return
	} else {
		PrintInOrder(n.left)
		fmt.Printf("%d ", n.key)
		PrintInOrder(n.right)
	}
}


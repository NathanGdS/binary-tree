package main

import (
	"fmt"

	"github.com/m1gwings/treedrawer/tree"
)

type Node struct {
	left  *Node
	right *Node
	value int
}

func (n *Node) insert(value int) {
	if n.value == 0 {
		n.insertFirstValue(value)
		return
	}

	if value < n.value {
		n.left = &Node{left: nil, right: nil, value: value}
		return
	}

	if value > n.value {
		n.right = &Node{left: nil, right: nil, value: value}
		return
	}

}

func (n *Node) insertFirstValue(value int) {
	n.value = value
}

func main() {
	var node Node
	node.insert(5)
	node.insert(4)
	node.insert(8)

	t := tree.NewTree(tree.NodeInt64(node.value))
	t.AddChild(tree.NodeInt64(node.left.value))
	t.AddChild(tree.NodeInt64(node.right.value))

	fmt.Println(t)
}

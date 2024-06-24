package main

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
}

func (n *Node) insertFirstValue(value int) {
	n.value = value
}

func main() {
	var node Node
	node.insert(2)
}

package node

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (n Node) Print() {
	fmt.Printf("%d ", n.Value)
}

func (n *Node) SetValue(value int) {
	n.Value = value
}

func CreatNode(v int) *Node {
	return &Node{Value: v}
}


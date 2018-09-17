package node

import "fmt"

func (n *Node) Travels() {
	//if n == nil {
	//	return
	//}
	//n.Left.Travels()
	//n.Print()
	//n.Right.Travels()
	n.TravelsFunc(func(node *Node) {
		node.Print()
		fmt.Println("hello")
	})
}

func (n *Node) TravelsFunc(f func(*Node)) {
	if n == nil {
		return
	}
	n.Left.TravelsFunc(f)
	f(n)
	n.Right.TravelsFunc(f)
}

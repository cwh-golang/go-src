package main

import (
	"fmt"
	"go-src/treerc/tree/node"
)

type myTreeNode struct {
	//对node使用组合来进行封装
	mynode *node.Node
}

func (mnode *myTreeNode) postOrder() {
	if mnode == nil || mnode.mynode == nil {
		return
	}
	left := myTreeNode{mnode.mynode.Left}
	right := myTreeNode{mnode.mynode.Right}
	left.postOrder()
	right.postOrder()
	mnode.mynode.Print()
}

func main() {
	n := node.Node{15, nil, nil}
	n.Left = node.CreatNode(30)
	n.Right = node.CreatNode(56)
	n.Travels()
	NodeCnt:=0
	n.TravelsFunc(func(i *node.Node) {
		NodeCnt++
	})
	fmt.Println(NodeCnt)
	//n.Print()
	//n.SetValue(45)
	//n.Print()
	fmt.Println()
	mn := myTreeNode{&n}
	mn.postOrder()
}

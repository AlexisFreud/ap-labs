package main

import (
	"sync"
)

// Node is the class to make the map
type Node struct {
	left         *Node
	right        *Node
	back         *Node
	front        *Node
	nextHop      *Node
	isExit       bool
	isBorder     bool
	stepsToExit int
	mux          *sync.Mutex
}

// NewNode creates nodes
func NewNode(isExit bool, isBorder bool, stepsToExit int) *Node {

	auxNode := &Node{}

	auxNode.left = nil
	auxNode.right = nil
	auxNode.back = nil
	auxNode.front = nil
	auxNode.nextHop = nil
	auxNode.isExit = isExit
	auxNode.isBorder = isBorder
	auxNode.stepsToExit = stepsToExit
	auxNode.mux = new(sync.Mutex)

	return auxNode
}

func (n *Node) getNext() []*Node{
	array := make([]*Node, 4)
	array[0] = n.front
	array[1] = n.right
	array[2] = n.back
	array[3] = n.left
	return array
}


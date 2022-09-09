package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}
type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	if node.nextNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}

func main() {
	list := &LinkedList{}
	list.AddToHead(1)
	list.AddToHead(2)
	fmt.Println(list.headNode.property)
	fmt.Println(list.headNode.nextNode)
}

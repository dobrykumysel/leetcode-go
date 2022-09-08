package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

// LinkedList represents a linked list
type MyLinkedList struct {
	head *Node
	len  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}
func (l *MyLinkedList) Insert(val int) {
	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

// Print displays all the nodes from linked list
func (l *MyLinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println(l.len)
		fmt.Println("Node: ", ptr)
		ptr = ptr.next
	}
}
func main() {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtHead(2)
	obj.AddAtHead(3)

	fmt.Printf("first: %v\n", obj.Get(1))
	fmt.Printf("second: %v\n", obj.Get(2))
	fmt.Printf("third: %v\n", obj.Get(3))

	obj.DeleteAtIndex(1)
	fmt.Printf("index : %v\n", obj.Get(1))
}

func (this *MyLinkedList) Get(index int) int {
	ptr := this.head
	for i := 1; i <= this.len; i++ {
		if i == index {
			return ptr.value
		}
		ptr = ptr.next
	}
	return -1
}

// добавка начала

func (this *MyLinkedList) AddAtHead(val int) {
	n := Node{}
	n.value = val
	n.next = this.head
	if this.len == 0 {
		this.head = &n
		this.len++
		return
	}
	this.head = &n
	this.len++
}

// Разбор В Хвост
func (this *MyLinkedList) AddAtTail(val int) {
	n := Node{}
	n.value = val
	if this.len == 0 {
		this.head = &n
		this.len++
		return
	}
	ptr := this.head
	for i := 0; i < this.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			this.len++
			return
		}
		ptr = ptr.next
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	ptr := this.head
	for i := 1; i <= this.len; i++ {
		if i == index {
			ptr.value = val
		}
		ptr = ptr.next
	}
}
func (this *MyLinkedList) GetNode(index int) Node {
	ptr := this.head
	for i := 1; i <= this.len; i++ {
		if i == index {
			return *ptr
		}
		ptr = ptr.next
	}
	return Node{}
}

// удалить ноду и изменить порядок
func (this *MyLinkedList) DeleteAtIndex(index int) {
	prevNode := this.GetNode(index - 1)
	prevNode.next = this.GetNode(index).next
	this.len--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

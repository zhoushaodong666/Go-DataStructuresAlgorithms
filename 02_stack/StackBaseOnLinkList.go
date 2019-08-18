package _2_stack

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkListStack struct {
	topNode *Node
}

func InitLinkListStack() *LinkListStack {
	return &LinkListStack{nil}
}

func (this *LinkListStack) IsEmpty() bool {
	return this.topNode == nil
}

func (this *LinkListStack) Push(v interface{}) {
	this.topNode = &Node{v, this.topNode}
}

func (this *LinkListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	val := this.topNode.data
	this.topNode = this.topNode.next
	return val
}

func (this *LinkListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.topNode.data
}

func (this *LinkListStack) Flush() {
	this.topNode = nil
}

func (this *LinkListStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty stack")
	}
	cur := this.topNode
	for cur != nil {
		fmt.Println(cur.data)
		cur = cur.next
	}
}

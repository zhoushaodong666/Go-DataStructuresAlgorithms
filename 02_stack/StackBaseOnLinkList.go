package _2_stack

import "fmt"

// 定义链表节点结构体
type Node struct {
	data interface{}
	next *Node
}

// 定义链栈结构体
type LinkListStack struct {
	topNode *Node
}

// 初始化链栈
func InitLinkListStack() *LinkListStack {
	return &LinkListStack{nil}
}

// 判断链栈是否为空
func (this *LinkListStack) IsEmpty() bool {
	return this.topNode == nil
}

// 入栈
func (this *LinkListStack) Push(v interface{}) {
	this.topNode = &Node{v, this.topNode}
}

// 出栈
func (this *LinkListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	val := this.topNode.data
	this.topNode = this.topNode.next
	return val
}

// 获取栈顶元素值
func (this *LinkListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.topNode.data
}

// 链栈置空
func (this *LinkListStack) Flush() {
	this.topNode = nil
}

// 打印链栈
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

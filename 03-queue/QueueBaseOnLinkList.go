package _3_queue

import "fmt"

// 定义链表节点
type Node struct {
	data interface{}
	next *Node
}

// 定义链表队列结构体
type LinkListQueue struct {
	front  *Node
	rear   *Node
	length uint
}

// 初始化链表队列
func InitLinkListQueue() *LinkListQueue {
	return &LinkListQueue{nil, nil, 0}
}

// 判断循环链表队列是否为空
func (this *LinkListQueue) IsEmpty() bool {
	return this.length == 0
}

// 入列
func (this *LinkListQueue) EnQueue(v interface{}) {
	node := &Node{v, nil}
	if this.rear == nil {
		this.front = node
	} else {
		this.rear.next = node
	}
	this.rear = node
	this.length++
}

func (this *LinkListQueue) DeQueue() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.front.data
	this.front = this.front.next
	this.length--
	return v
}

// 返回队头元素
func (this *LinkListQueue) Front() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.front.data
}

// 清空链表
func (this *LinkListQueue) Flush() {
	this.front = nil
	this.rear = nil
	this.length = 0
}

// 打印队列
func (this *LinkListQueue) Print() {
	if this.IsEmpty() {
		fmt.Println("empty queue")
		return
	}

	formateStr := "head"
	cur := this.front
	for ; cur != nil; cur = cur.next {
		formateStr += fmt.Sprintf("<-%+v", cur.data)
	}
	formateStr += "<-tail\n"
	fmt.Print(formateStr)
}

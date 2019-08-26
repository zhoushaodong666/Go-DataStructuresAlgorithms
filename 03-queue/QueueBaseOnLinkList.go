package _3_queue

// 定义链表节点
type Node struct {
	data interface{}
	next *Node
}

// 定义链表队列结构体
type LinkListQueue struct {
	head   *Node
	length uint
}

// 初始化链表队列
func InitLinkListQueue() *LinkListQueue {
	return &LinkListQueue{
		head:   &Node{nil, nil},
		length: 0,
	}
}

// 判断循环链表队列是否为空
func (this *LinkListQueue) IsEmpty() bool {
	return this.length == 0
}

// 入列
func (this *LinkListQueue) EnQueue(v interface{}) bool {
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	newNode := &Node{v, nil}
	cur.next = newNode
	this.length++
	return true
}

func (this *LinkListQueue) DeQueue(v interface{}) (interface{}, error) {
	if this.IsEmpty() {
		return nil, nil
	}

	// 虚拟头节点
	var pre *Node = nil
	// 第一个有效节点
	cur := this.head.next

	for cur != nil {
		pre = cur
		cur = cur.next
	}

	if pre != nil {
		this.head = cur
		this.length--
		return nil, nil
	}
	return nil, nil
}

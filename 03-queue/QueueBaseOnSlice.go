package _3_queue

import "fmt"

// 定义顺序循环队列结构体
type SeqSliceQueue struct {
	// 数据集
	data []interface{}
	// 队头指针
	front int
	// 队尾指针
	rear int
	// 队列长度
	cap int
}

// 初始化循环队列 传入n为 队列长度
func InitSeqSliceQueue(size int) *SeqSliceQueue {
	return &SeqSliceQueue{make([]interface{}, size), 0, 0, size}
}

// 队列是否为空
func (this *SeqSliceQueue) IsEmpty() bool {
	if this.front == this.rear {
		return true
	}
	return false
}

// 队列是否已满
func (this *SeqSliceQueue) IsFull() bool {
	return (this.rear+1)%this.cap == this.front
}

// 入队
func (this *SeqSliceQueue) EnQueue(v interface{}) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.rear] = v
	this.rear = (this.rear + 1) % this.cap
	return true
}

// 出队
func (this *SeqSliceQueue) DeQueue() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.data[this.front]
	this.front = (this.front + 1) % this.cap
	return v
}

// 返回队列内元素个数
func (this *SeqSliceQueue) Length() int {
	return (this.rear - this.front + this.cap) % this.cap
}

// 返回队列首个元素
func (this *SeqSliceQueue) Front() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.data[this.front]
	return v
}

// 清空队列
func (this *SeqSliceQueue) Flush() {
	this.front = 0
	this.rear = 0
}

func (this *SeqSliceQueue) Print() {
	if this.IsEmpty() {
		fmt.Println("empty queue")
		return
	}
	fmt.Print("队头")
	for i := this.front; i != this.rear; {
		fmt.Print(this.data[i], " ")
		i = (i + 1) % this.cap
	}
	fmt.Print("队尾")
	fmt.Println()

}

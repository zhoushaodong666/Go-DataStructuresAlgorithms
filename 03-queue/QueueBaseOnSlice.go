package _3_queue

import "fmt"

// 定义顺序循环队列结构体
type SliceQueue struct {
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
func InitSeqSliceQueue(size int) *SliceQueue {
	return &SliceQueue{make([]interface{}, size), 0, 0, size}
}

// 队列是否为空
func (this *SliceQueue) IsEmpty() bool {
	return this.front == this.rear
}

// 队列是否已满
func (this *SliceQueue) IsFull() bool {
	return this.rear == this.cap
}

// 入队
func (this *SliceQueue) EnQueue(v interface{}) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.rear] = v
	this.rear++
	return true
}

// 出队
func (this *SliceQueue) DeQueue() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.data[this.front]
	this.front++
	return v
}

// 返回队列内元素个数
func (this *SliceQueue) Length() int {
	return this.rear - this.front
}

// 返回队列首个元素
func (this *SliceQueue) Front() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.data[this.front]
	return v
}

// 清空队列
func (this *SliceQueue) Flush() {
	this.front = 0
	this.rear = 0
}

func (this *SliceQueue) Print() {
	if this.IsEmpty() {
		fmt.Println("empty queue")
		return
	}
	fmt.Print("队头")
	for i := this.front; i != this.rear; i++ {
		fmt.Print(this.data[i], " ")
	}
	fmt.Print("队尾")
	fmt.Println()

}

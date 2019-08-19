package _3_queue

import (
	"fmt"
	"testing"
)

func TestSeqSliceQueue_EnQueue(t *testing.T) {
	q := InitSeqSliceQueue(5)
	for i := 0; i < 5; i++ {
		fmt.Println(q.EnQueue(i + 1))
	}
	q.Print()
}

func TestSeqSliceQueue_DeQueue(t *testing.T) {
	q := InitSeqSliceQueue(5)
	for i := 0; i < 5; i++ {
		fmt.Println(q.EnQueue(i + 1))
	}
	q.Print()
	for i := 0; i < 5; i++ {
		fmt.Println(q.DeQueue())
		q.Print()
	}
}

func TestSeqSliceQueue_Front(t *testing.T) {
	q := InitSeqSliceQueue(5)
	for i := 0; i < 5; i++ {
		fmt.Println(q.EnQueue(i + 1))
	}
	q.Print()
	for i := 0; i < 5; i++ {
		q.DeQueue()
		q.Print()
		fmt.Println("队头元素为：", q.Front())
	}
}

func TestSeqSliceQueue_IsEmpty(t *testing.T) {
	q := InitSeqSliceQueue(5)
	q.EnQueue(1)
	q.Print()
	fmt.Println("队列是否为空：", q.IsEmpty())

	q.DeQueue()
	q.Print()
	fmt.Println("队列是否为空：", q.IsEmpty())

	q.DeQueue()
	q.Print()
	fmt.Println("队列是否为空：", q.IsEmpty())

}

func TestSeqSliceQueue_IsFull(t *testing.T) {
	q := InitSeqSliceQueue(5)
	for i := 0; i < 5; i++ {
		fmt.Println(q.EnQueue(i + 1))
	}
	q.Print()
	fmt.Println("队列是否已满：", q.IsFull())

	q.DeQueue()
	q.Print()
	fmt.Println("队列是否已满：", q.IsFull())

	q.DeQueue()
	q.Print()
	fmt.Println("队列是否已满：", q.IsFull())

	q.EnQueue(5)
	q.Print()
	fmt.Println("队列是否已满：", q.IsFull())

	q.EnQueue(6)
	q.Print()
	fmt.Println("队列是否已满：", q.IsFull())

}

func TestSeqSliceQueue_Length(t *testing.T) {
	q := InitSeqSliceQueue(5)
	for i := 0; i < 5; i++ {
		q.EnQueue(i + 1)
	}
	q.Print()
	fmt.Println("队列长度为：", q.Length())

	for i := 0; i < 5; i++ {
		q.DeQueue()
		q.Print()
		fmt.Println("队列长度为：", q.Length())
	}

}

func TestSeqSliceQueue_Flush(t *testing.T) {
	q := InitSeqSliceQueue(5)
	for i := 0; i < 5; i++ {
		q.EnQueue(i + 1)
	}
	q.Print()

	q.Flush()
	q.Print()
}

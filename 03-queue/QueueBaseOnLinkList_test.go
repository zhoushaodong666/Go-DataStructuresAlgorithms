package _3_queue

import (
	"fmt"
	"testing"
)

func TestLinkListQueue_EnQueue(t *testing.T) {
	q := InitLinkListQueue()
	for i := 0; i < 5; i++ {
		q.EnQueue(i + 1)
	}
	q.Print()

}

func TestLinkListQueue_DeQueue(t *testing.T) {
	q := InitLinkListQueue()
	for i := 0; i < 5; i++ {
		q.EnQueue(i + 1)
	}
	q.Print()
	for i := 0; i < 6; i++ {
		fmt.Println("出列元素：", q.DeQueue())
		fmt.Print("当前队列为：")
		q.Print()
	}
	q.Print()
}

func TestLinkListQueue_Front(t *testing.T) {
	q := InitLinkListQueue()
	for i := 0; i < 5; i++ {
		q.EnQueue(i + 1)
	}
	fmt.Println("队头元素为：", q.Front())
	for i := 0; i < 6; i++ {
		q.DeQueue()
		fmt.Print("当前队列为：")
		q.Print()
		fmt.Println("队头元素为：", q.Front())
	}
}

func TestLinkListQueue_IsEmpty(t *testing.T) {
	q := InitLinkListQueue()
	q.EnQueue(1)
	q.Print()
	fmt.Println("当前队列是否为空:", q.IsEmpty())
	q.DeQueue()
	q.Print()
	fmt.Println("当前队列是否为空:", q.IsEmpty())
}

func TestLinkListQueue_Flush(t *testing.T) {
	q := InitLinkListQueue()
	for i := 0; i < 5; i++ {
		q.EnQueue(i + 1)
	}
	q.Print()
	q.Flush()
	q.Print()
	for i := 5; i < 10; i++ {
		q.EnQueue(i + 1)
	}
	q.Print()
	q.DeQueue()
	q.Print()
}

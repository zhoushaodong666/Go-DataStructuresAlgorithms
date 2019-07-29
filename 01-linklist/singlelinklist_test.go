package _1_linklist

import (
	"fmt"
	"testing"
)

/*
Project: 单链表测试
Author: zhoushaodong
Email: 1138894663@qq.com
Github: https://github.com/zhoushaodong666
Date: 2019/7/23
Time: 15:46
*/

func TestInsertToHead(t *testing.T) {
	l := InitLinkList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i + 1)
	}

	l.Print()
	t.Log("linklist length = ", l.length)
}

func TestInsertToTail(t *testing.T) {
	l := InitLinkList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	t.Log("linklist length = ", l.length)
}

func TestInsertToBefore(t *testing.T) {
	l := InitLinkList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	node0 := l.FindByIndex(0)
	node1 := l.FindByIndex(9)
	l.InsertToBefore(node0, 11)
	l.InsertToBefore(node1, 12)
	l.Print()

}

func TestDeleteHead(t *testing.T) {
	l := InitLinkList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	fmt.Println("节点数：", l.length)

	t.Log("开始删除节点")
	for j := 0; j < 10; j++ {
		l.DeleteHead()
		l.Print()
	}
}

func TestDeleteTail(t *testing.T) {
	l := InitLinkList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	t.Log("节点数：", l.length)

	t.Log("开始删除节点")
	for j := 0; j < 10; j++ {
		l.DeleteTail()
		l.Print()
	}
}

func TestFindByIndex(t *testing.T) {
	l := InitLinkList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(5))
	t.Log(l.FindByIndex(2))
	t.Log(l.FindByIndex(10))
}

func TestDeleteNode(t *testing.T) {
	l := InitLinkList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	t.Log(l.DeleteNode(l.head.next))
	l.Print()

	t.Log(l.DeleteNode(l.head.next.next))
	l.Print()
}

package _1_linklist

import (
	"fmt"
	"testing"
)

/*
Project: 双向链表测试
Author: zhoushaodong
Email: 1138894663@qq.com
Github: https://github.com/zhoushaodong666
Date: 2019/7/26
Time: 14:09
*/

func TestDLinkList_InsertToHead(t *testing.T) {
	l := InitDLinkList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i + 1)
	}
	l.Print()
	t.Log("linklist length = ", l.length)
}

func TestDLinkList_InsertToTail(t *testing.T) {
	l := InitDLinkList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	t.Log("linklist length = ", l.length)
}

func TestDLinkList_DeleteHead(t *testing.T) {
	l := InitDLinkList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	for i := 0; i < 10; i++ {
		l.DeleteHead()
		l.Print()
		fmt.Println("link list length = ", l.length)
	}
}

func TestDLinkList_DeleteTail(t *testing.T) {
	l := InitDLinkList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	for i := 0; i < 10; i++ {
		l.DeleteTail()
		l.Print()
		fmt.Println("link list length = ", l.length)
	}
}

func TestDLinkList_FindByIndex(t *testing.T) {
	l := InitDLinkList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	node0 := l.FindByIndex(0)
	node1 := l.FindByIndex(1)
	node2 := l.FindByIndex(5)
	node3 := l.FindByIndex(10)
	node4 := l.FindByIndex(11)
	fmt.Println(node0, node1, node2, node3, node4)
}

package _1_linklist

import "testing"

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

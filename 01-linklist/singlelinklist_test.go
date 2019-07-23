package _1_linklist

import "testing"

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
	println("linklist length = ", l.length)
}

func TestInsertToTail(t *testing.T) {

	l := InitLinkList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	println("linklist length = ", l.length)
}

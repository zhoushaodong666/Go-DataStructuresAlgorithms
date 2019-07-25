package _1_linklist

import (
	"fmt"
	"github.com/pkg/errors"
)

/*
Project: 单链表的基本操作
Author: zhoushaodong
Email: 1138894663@qq.com
Github: https://github.com/zhoushaodong666
Date: 2019/7/23
Time: 15:30
*/

//单链表节点定义
type ListNode struct {
	value interface{}
	next  *ListNode
}

//单链表定义
type LinkList struct {
	head   *ListNode
	length uint
}

//初始化单链表
func InitLinkList() *LinkList {
	linklist := &LinkList{NewListNode(nil), 0}
	if nil != linklist {
		return linklist
	}
	return nil
}

//返回一个新节点
func NewListNode(v interface{}) *ListNode {
	node := &ListNode{v, nil}
	if nil != node {
		return node
	}
	return nil
}

//获取节点值
func (this *ListNode) GetValue() interface{} {
	return this.value
}

//判断链表是否为空
func (this *LinkList) IsEmpty() bool {
	return this.length == 0
}

//将新节点插入链表头部
func (this *LinkList) InsertToHead(v interface{}) (bool, error) {
	return this.InsertToAfter(this.head, v)
}

//将新节点插入链表尾部
func (this *LinkList) InsertToTail(v interface{}) (bool, error) {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertToAfter(cur, v)
}

//在某个节点后面插入节点
func (this *LinkList) InsertToAfter(p *ListNode, v interface{}) (bool, error) {
	if nil == p {
		return false, errors.New("p node is nil")
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++
	return true, nil
}

//在某个节点前面插入节点
func (this *LinkList) InsertToBefore(p *ListNode, v interface{}) (bool, error) {
	if nil == p || p == this.head {
		return false, errors.New("p node is nil or is head node")
	}
	//第一个有效节点
	cur := this.head.next
	//虚拟头结点
	pre := this.head
	for nil != cur {
		//找到指定的节点 跳出循环
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	//循环结束还没有找到指定的节点
	if nil == cur {
		return false, errors.New("p node is not in linklist")
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true, nil
}

//删除传入的节点
func (this *LinkList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		//找到需要删除的节点，跳出循环
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	//循环结束还未找到传入的节点
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	this.length--
	return true
}

//删除头部节点
func (this *LinkList) DeleteHead() bool {
	//链表为空
	if this.IsEmpty() {
		return true
	}
	return this.DeleteNode(this.head.next)
}

//删除尾部节点
func (this *LinkList) DeleteTail() bool {
	if this.IsEmpty() {
		return true
	}

	pre := this.head
	cur := this.head.next
	for nil != cur.next {
		pre = cur
		cur = cur.next
	}
	pre.next = nil
	this.length--
	return true
}

//通过索引查找节点 索引从0开始
func (this *LinkList) FindByIndex(index uint) *ListNode {
	//空链或索引超过链表长度就返回nil
	if this.IsEmpty() || index >= this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//反转单链表

func (this *LinkList) ReverseLinkList() {
	var pre *ListNode = nil
	cur := this.head.next
	for nil != cur {
		tempNextNode := cur.next
		cur.next = pre
		pre = cur
		cur = tempNextNode
	}
	this.head.next = pre
}

//打印链表
func (this *LinkList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	if format != "" {
		format += "->nil"
	} else {
		format = "nil"
	}
	fmt.Println(format)
}

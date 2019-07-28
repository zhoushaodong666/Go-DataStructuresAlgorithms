package _1_linklist

import (
	"errors"
	"fmt"
)

/*
Project:
Author: zhoushaodong
Email: 1138894663@qq.com
Github: https://github.com/zhoushaodong666
Date: 2019/7/26
Time: 14:08
*/

//定义双向链表结构体
type DLinkList struct {
	head   *DListNode
	length uint
}

//定义双向链表节点结构体
type DListNode struct {
	value interface{}
	pre   *DListNode
	next  *DListNode
}

//初始化单链表
func InitDLinkList() *DLinkList {
	linklist := &DLinkList{NewDListNode(nil), 0}
	if nil != linklist {
		return linklist
	}
	return nil
}

//返回一个新节点
func NewDListNode(v interface{}) *DListNode {
	node := &DListNode{v, nil, nil}
	if nil != node {
		return node
	}
	return nil
}

//获取节点值
func (this *DListNode) GetValue() interface{} {
	return this.value
}

//判断链表是否为空
func (this *DLinkList) IsEmpty() bool {
	return this.length == 0
}

//将新节点插入链表头部
func (this *DLinkList) InsertToHead(v interface{}) (bool, error) {
	return this.InsertToAfter(this.head, v)
}

//将新节点插入链表尾部
func (this *DLinkList) InsertToTail(v interface{}) (bool, error) {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertToAfter(cur, v)
}

//在某个节点后面插入节点
func (this *DLinkList) InsertToAfter(p *DListNode, v interface{}) (bool, error) {

	if nil == p {
		return false, errors.New("p node is nil")
	}
	newNode := NewDListNode(v)

	//当尾部插入时，最后一个节点的next域为nil 所以不用next节点的pre域建立联系
	if nil != p.next {
		p.next.pre = newNode
	}
	//新节点与后继节点建立连接
	newNode.next = p.next

	//新节点与前驱节点建立连接
	newNode.pre = p
	p.next = newNode

	//lenght自增
	this.length++
	return true, nil
}

//在某个节点前面插入节点
func (this *DLinkList) InsertToBefore(p *DListNode, v interface{}) (bool, error) {
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
	newNode := NewDListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true, nil
}

//删除传入的节点 双向链表可以很方便的通过pre域找到前驱 所以不需要遍历
func (this *DLinkList) DeleteNode(p *DListNode) bool {
	if nil == p || this.IsEmpty() {
		return false
	}
	//前驱节点
	pre := p.pre
	//后继节点
	next := p.next
	//有后继节点时和前驱节点建立联系
	if p.next != nil {
		next.pre = pre
	}
	//前驱和后继节点建立联系
	pre.next = next

	this.length--
	return true
}

//删除头部节点
func (this *DLinkList) DeleteHead() bool {
	return this.DeleteNode(this.head.next)
}

//删除尾部节点
func (this *DLinkList) DeleteTail() bool {
	cur := this.head.next
	for nil != cur.next {
		cur = cur.next
	}
	return this.DeleteNode(cur)
}

//通过索引查找节点 索引从0开始
func (this *DLinkList) FindByIndex(index uint) *DListNode {
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

//打印链表
func (this *DLinkList) Print() {
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

package _2_stack

import "fmt"

/*
Project: 顺序栈的基本操作(基于切片)
Author: zhoushaodong
Email: 1138894663@qq.com
Github: https://github.com/zhoushaodong666
Date: 2019/7/29
Time: 15:02
*/

type SliceStack struct {
	//数据
	data []interface{}
	//栈顶指针索引
	top int
	Stack
}

//初始化顺序栈
func InitSliceStack() *SliceStack {
	return &SliceStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (this *SliceStack) IsEmpty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

//压栈(入栈)
func (this *SliceStack) Push(v interface{}) {
	//当栈为空时
	if this.IsEmpty() {
		this.top = 0
	} else {
		this.top += 1
	}

	//当栈空间不足时 使用append函数插入 append会内部进行扩容
	if this.top > len(this.data)-1 {
		this.data = append(this.data, v)
	} else {
		this.data[this.top] = v
	}
}

//出栈
func (this *SliceStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	value := this.data[this.top]
	this.top--
	return value
}

//获取栈顶元素
func (this *SliceStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.data[this.top]
}

//重置栈
func (this *SliceStack) Flush() {
	this.top = -1
}

//打印栈
func (this *SliceStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty statck")
	} else {
		for i := this.top; i >= 0; i-- {
			fmt.Println(this.data[i])
		}
	}
}

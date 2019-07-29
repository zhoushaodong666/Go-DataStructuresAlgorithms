package _2_stack

import (
	"fmt"
	"testing"
)

/*
Project:
Author: zhoushaodong
Email: 1138894663@qq.com
Github: https://github.com/zhoushaodong666
Date: 2019/7/29
Time: 15:52
*/

func TestSliceStack_Push(t *testing.T) {
	s := InitSliceStack()
	for i := 0; i < 20; i++ {
		s.Push(i + 1)
	}
	s.Print()
}

func TestSliceStack_Pop(t *testing.T) {
	s := InitSliceStack()
	for i := 0; i < 20; i++ {
		s.Push(i + 1)
	}
	s.Print()
	fmt.Println("----------")
	p1 := s.Pop()
	fmt.Println("出栈元素值为", p1)
	fmt.Println("----------")
	s.Print()
}

func TestSliceStack_IsEmpty(t *testing.T) {
	s := InitSliceStack()
	fmt.Println("初始化后的栈空间:", s.IsEmpty())
	for i := 0; i < 5; i++ {
		s.Push(i + 1)
	}
	s.Print()
	fmt.Println("压栈后栈空间:", s.IsEmpty())
	for i := 0; i < 5; i++ {
		p1 := s.Pop()
		fmt.Println("-----")
		fmt.Println("出栈元素值为", p1)
		fmt.Println("-----")
		s.Print()
		fmt.Println("出栈后栈空间:", s.IsEmpty())
	}
}

func TestSliceStack_Top(t *testing.T) {
	s := InitSliceStack()
	for i := 0; i < 5; i++ {
		s.Push(i + 1)
	}
	s.Print()
	top := s.Top()
	fmt.Println("栈顶元素值：", top)
}

func TestSliceStack_Flush(t *testing.T) {
	s := InitSliceStack()
	for i := 0; i < 5; i++ {
		s.Push(i + 1)
	}
	s.Print()
	s.Flush()
	fmt.Println("----------")
	s.Print()
}

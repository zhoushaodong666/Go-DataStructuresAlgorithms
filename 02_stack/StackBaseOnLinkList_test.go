package _2_stack

import (
	"fmt"
	"testing"
)

func TestLinkListStack_Push(t *testing.T) {
	s := InitLinkListStack()
	for i := 0; i < 20; i++ {
		s.Push(i + 1)
	}
	s.Print()
}

func TestLinkListStack_Pop(t *testing.T) {
	s := InitLinkListStack()
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

func TestLinkListStack_IsEmpty(t *testing.T) {
	s := InitLinkListStack()
	fmt.Println("初始化后的栈空间是否为空:", s.IsEmpty())
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

func TestLinkListStack_Top(t *testing.T) {
	s := InitLinkListStack()
	for i := 0; i < 5; i++ {
		s.Push(i + 1)
	}
	s.Print()
	top := s.Top()
	fmt.Println("栈顶元素值：", top)
}

func TestLinkListStack_Flush(t *testing.T) {
	s := InitLinkListStack()
	for i := 0; i < 5; i++ {
		s.Push(i + 1)
	}
	s.Print()
	s.Flush()
	fmt.Println("----------")
	s.Print()
}

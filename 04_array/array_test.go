package _4_array

import (
	"fmt"
	"testing"
)

func TestArray_Insert(t *testing.T) {
	arr := InitArray(5)
	for i := 0; i < 5; i++ {
		arr.Insert(uint(i), i+1)
	}
	arr.Print()
}

func TestArray_IsIndexOutOfRange(t *testing.T) {
	arr := InitArray(5)
	for i := 0; i < 5; i++ {
		arr.Insert(uint(i), i+1)
	}
	arr.Print()
	fmt.Println("是否下标越界：", arr.IsIndexOutOfRange(5))

}

func TestArray_Len(t *testing.T) {
	arr := InitArray(5)
	for i := 0; i < 5; i++ {
		arr.Insert(uint(i), i+1)
		arr.Print()
		fmt.Println("数组长度：", arr.Len())
	}
}

func TestArray_InsertToTail(t *testing.T) {
	arr := InitArray(5)
	for i := 0; i < 5; i++ {
		arr.InsertToTail(i + 1)
	}
	arr.Print()
}

func TestArray_Delete(t *testing.T) {
	arr := InitArray(5)
	for i := 0; i < 5; i++ {
		arr.InsertToTail(i + 1)
	}
	arr.Print()

	arr.Delete(0)
	arr.Print()

	arr.Delete(0)
	arr.Print()

	arr.Delete(5)
	arr.Print()

}

func TestArray_Find(t *testing.T) {
	arr := InitArray(5)
	for i := 0; i < 5; i++ {
		arr.InsertToTail(i + 1)
	}
	arr.Print()

	v, _ := arr.Find(0)
	fmt.Println(v)

	v, _ = arr.Find(1)
	fmt.Println(v)

	v, _ = arr.Find(2)
	fmt.Println(v)

	v, _ = arr.Find(3)
	fmt.Println(v)

	v, _ = arr.Find(4)
	fmt.Println(v)

	v, err := arr.Find(5)
	fmt.Println(v)
	fmt.Println(err)
}

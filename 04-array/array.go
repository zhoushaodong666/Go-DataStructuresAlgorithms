package _4_array

import (
	"errors"
	"fmt"
)

// 定义数组结构体
type Array struct {
	data   []interface{}
	length uint
}

// 初始化数组
func InitArray(size uint) *Array {
	if size <= 0 {
		return nil
	}
	return &Array{
		data:   make([]interface{}, size, size),
		length: 0,
	}
}

// 获取数组长度
func (this *Array) Len() uint {
	return this.length
}

// 判断索引是否越界
func (this *Array) IsIndexOutOfRange(index uint) bool {
	if index >= uint(cap(this.data)) {
		return true
	}
	return false
}

// 根据索引返回元素
func (this *Array) Find(index uint) (interface{}, error) {
	if this.IsIndexOutOfRange(index) {
		return nil, errors.New("index out Of range")
	}
	return this.data[index], nil
}

// 根据传入索引插入元素
func (this *Array) Insert(index uint, v interface{}) error {
	if this.Len() == uint(cap(this.data)) {
		return errors.New("full array")
	}

	if this.IsIndexOutOfRange(index) {
		return errors.New("index out of range")
	}

	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	return nil
}

// 将元素插入数组末尾
func (this *Array) InsertToTail(v int) error {
	return this.Insert(this.Len(), v)
}

// 根据索引删除元素
func (this *Array) Delete(index uint) (interface{}, error) {
	if this.IsIndexOutOfRange(index) {
		return 0, errors.New("index out of range")
	}
	v := this.data[index]
	for i := index; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v, nil
}

// 打印数组
func (this *Array) Print() {
	var format string
	for i := uint(0); i < this.Len(); i++ {
		format += fmt.Sprintf(" %+v", this.data[i])
	}
	fmt.Println(format)
}

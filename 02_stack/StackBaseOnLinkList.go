package _2_stack

type NodeStack struct {
	data interface{}
	next *NodeStack
}

type LinkListStack struct {
	head *NodeStack
	top  int
}

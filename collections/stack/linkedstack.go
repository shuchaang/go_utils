package stack

import (
	"errors"
)

type linkedStack struct {
	size int
	head *node
	tail *node
}

type node struct {
	value interface{}
	next  *node
	prev  *node
}

func NewLinkedStack() *linkedStack {
	head := &node{
		value: nil,
		next:  nil,
		prev:  nil,
	}
	tail := &node{
		value: nil,
		next:  nil,
		prev:  nil,
	}
	head.next = tail
	tail.prev = head
	return &linkedStack{
		size: 0,
		head: head,
		tail: tail,
	}
}

func (ls *linkedStack) Push(v interface{}) {
	newNode := &node{
		value: v,
		next:  nil,
		prev:  nil,
	}

	p := ls.tail.prev
	p.next = newNode
	newNode.prev = p
	newNode.next = ls.tail
	ls.tail.prev = newNode
	ls.size++
}

func (ls *linkedStack) Pop() interface{} {
	if ls.IsEmpty() {
		return errors.New("stack is empty")
	}
	end := ls.tail.prev.prev
	tmp := end.next.value
	end.next = ls.tail
	ls.tail.prev = end
	ls.size--
	return tmp
}

func (ls *linkedStack) Len() int {
	return ls.size
}
func (ls *linkedStack) Peek() interface{} {
	if ls.IsEmpty() {
		return errors.New("stack is empty")
	}
	return ls.tail.prev.value
}

func (ls *linkedStack) IsEmpty() bool {
	return ls.size == 0
}

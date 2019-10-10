package stack

import "testing"

func TestLinkedStack_Len(t *testing.T) {
	stack := NewLinkedStack()
	stack.Push(1)
	stack.Push(222)
	stack.Push("123")

	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())

}

package linkedqueue

import (
	"fmt"
	"testing"
)

func TestLinkedQueue_Offer(t *testing.T) {
	queue := NewLinkedQueue()
	fmt.Println(queue.tail)
}


func TestLinkedQueue_Poll(t *testing.T) {
	q:=NewLinkedQueue()
	q.Offer(1)
	q.Offer(2)
	t.Log(q.head.value)
	t.Log(q.Peek())
	t.Log(q.Size())
	t.Log(q.Poll())
	t.Log(q.Size())
}
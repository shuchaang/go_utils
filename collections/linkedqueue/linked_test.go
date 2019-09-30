package linkedqueue

import (
	"testing"
)

func TestLinkedQueue_Offer(t *testing.T) {
	q := NewLinkedQueue()
	q.Offer(11)
	q.Offer(22)
	q.Offer(33)
	t.Log(q.Peek())
	t.Log(q.Size())
	t.Log(q.Poll())
	t.Log(q.Size())
	t.Log(q.Peek())
}

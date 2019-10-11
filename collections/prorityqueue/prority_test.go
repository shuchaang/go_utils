package prorityqueue

import (
	"testing"
)

func TestPQ_Len(t *testing.T) {
	q := NewPriorityQ()
	q.Push(&Node{priority: 8, value: 1})
	q.Push(&Node{priority: 7, value: 2})
	q.Push(&Node{priority: 9, value: 3})

	top := q.Top()
	t.Log(top.priority, top.value)

	for q.Len() > 0 {
		x := q.Pop()
		t.Log(x.priority, x.value)
	}
}

func (this *Node) Less(node *Node) bool {
	return this.priority < node.priority
}

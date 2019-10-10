package prorityqueue

import (
	"container/heap"
	"testing"
)

func TestPQ_Len(t *testing.T) {

	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	pq := make(PQ, len(items))
	i := 0

	for value, priority := range items {
		pq[i] = &Node{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	t.Log(pq)

	item := &Node{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Node)
		t.Log(item.value)
		t.Log(item.priority)
	}

}

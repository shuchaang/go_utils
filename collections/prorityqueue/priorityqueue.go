package prorityqueue

import "container/heap"

type Node struct {
	value    interface{} //每个节点存储的元素
	priority int         //优先级
}

type Interface interface {
	Less(node *Node) bool
}

type sorter []Interface

// Implement heap.Interface: Push, Pop, Len, Less, Swap
func (s *sorter) Push(x interface{}) {
	*s = append(*s, x.(Interface))
}

func (s *sorter) Pop() interface{} {
	n := len(*s)
	if n > 0 {
		x := (*s)[n-1]
		*s = (*s)[0 : n-1]
		return x
	}
	return nil
}

func (s *sorter) Len() int {
	return len(*s)
}

func (s *sorter) Less(i, j int) bool {
	return (*s)[i].Less((*s)[j].(*Node))
}

func (s *sorter) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

// Define priority queue struct
type PriorityQueue struct {
	s *sorter
}

func NewPriorityQ() *PriorityQueue {
	q := &PriorityQueue{s: new(sorter)}
	heap.Init(q.s)
	return q
}

func (q *PriorityQueue) Push(node *Node) {
	heap.Push(q.s, node)
}

func (q *PriorityQueue) Pop() *Node {
	return heap.Pop(q.s).(*Node)
}

func (q *PriorityQueue) Top() *Node {
	if len(*q.s) > 0 {
		return (*q.s)[0].(*Node)
	}
	return nil
}

func (q *PriorityQueue) Fix(x Interface, i int) {
	(*q.s)[i] = x
	heap.Fix(q.s, i)
}

func (q *PriorityQueue) Remove(i int) Interface {
	return heap.Remove(q.s, i).(Interface)
}

func (q *PriorityQueue) Len() int {
	return q.s.Len()
}

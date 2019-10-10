package prorityqueue

import "container/heap"

type Node struct {
	value    interface{} //每个节点存储的元素
	priority int         //优先级
	index    int         //堆中的索引
}

type PQ []*Node

// 下面的方法实现了 go/src/sort/sort.go:14 接口的方法

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 往 pq 中放 entry
func (pq *PQ) Push(x interface{}) {
	temp := x.(*Node)
	temp.index = len(*pq)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 entry
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	temp.index = -1 // for safety
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

// update modifies the priority and value of an entry in the queue.
func (pq *PQ) update(entry *Node, value string, priority int) {
	entry.value = value
	entry.priority = priority
	heap.Fix(pq, entry.index)
}

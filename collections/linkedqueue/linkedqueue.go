package linkedqueue

//链表队列
type linkedQueue struct {
	size int
	head *node
	tail *node
}

//链表Node结构
type node struct {
	value interface{}
	next  *node
	prev  *node
}

//----------------------- linked queue--------------------
// head <--> node<-->node<-->tail

func NewLinkedQueue() *linkedQueue {
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
	return &linkedQueue{
		size: 0,
		head: head,
		tail: tail,
	}
}

func (q *linkedQueue) Offer(v interface{}) {
	newNode := &node{
		value: v,
		next:  nil,
		prev:  nil,
	}
	if q.Size() == 0 {
		q.head.next = newNode
		q.tail.prev = newNode
		newNode.prev = q.head
		newNode.next = q.tail
	} else {
		newNode.prev = q.tail.prev
		q.tail.prev.next = newNode
		newNode.next = q.tail
		q.tail.prev = newNode
	}
	q.size++
}

func (q *linkedQueue) Size() int {
	return q.size
}

func (q *linkedQueue) Poll() interface{} {
	if q.Size() == 0 {
		return nil
	}
	tmp := q.head.next
	q.head.next = tmp.next
	tmp.next.prev = q.head
	data := tmp.value
	tmp = nil
	q.size--
	return data
}

func (q *linkedQueue) Peek() interface{} {
	return q.head.next.value
}

package collections


type Queue interface {
	//放入数据
	Offer(interface{})
	//取出数据
	Poll() interface{}
	//获取第一个元素 不出队
	Peek() interface{}
	//queue size
	Size() int
}


package collections

type Stack interface {
	Push(interface{})
	Pop() interface{}
	Len() int
	Peek() interface{}
	IsEmpty() bool
}

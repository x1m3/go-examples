package priorityqueues

type Queue interface {
	Push(something interface{}, priority int)
	Pop() interface{}
}

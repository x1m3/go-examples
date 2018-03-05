package linkedListWay

import "sync"

type item struct {
	value interface{}
	next  *item
}

func newItem(val interface{}) *item {
	return &item{value: val, next: nil}
}

type LinkedQueue struct {
	sync.RWMutex
	first *item
	last  *item
	Count int
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{first: nil, last: nil, Count: 0}
}

func (q *LinkedQueue) Push(something interface{}) {
	n := newItem(something)

	q.Lock()

	q.Count++

	if q.first == nil {
		q.first = n
		q.last = n
	} else {
		q.last.next = n
		q.last = n
	}

	q.Unlock()
}

func (q *LinkedQueue) Pop() interface{} {
	q.Lock()

	switch q.Count {
	case 0:
		q.Unlock()
		return nil
	case 1:
		r := q.first
		q.first = nil
		q.last = nil
		q.Count--
		q.Unlock()
		return r.value
	default:
		r := q.first
		q.first = q.first.next
		q.Count--
		q.Unlock()
		return r.value
	}
}

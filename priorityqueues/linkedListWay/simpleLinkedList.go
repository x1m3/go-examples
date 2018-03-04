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
	count int
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{first: nil, last: nil, count: 0}
}

func (q *LinkedQueue) Push(something interface{}) {
	n := newItem(something)

	q.Lock()
	defer q.Unlock()

	q.count++

	if q.first == nil {
		q.first = n
		q.last = n
	} else {
		q.last.next = n
		q.last = n
	}
}

func (q *LinkedQueue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()

	switch q.count {
	case 0:
		return nil
	case 1:
		r := q.first
		q.first = nil
		q.last = nil
		q.count--
		return r.value
	default:
		r := q.first
		q.first = q.first.next
		q.count--
		return r.value
	}
}

func (q *LinkedQueue) Count() int {
	q.RLock()
	defer q.RUnlock()

	return q.count
}

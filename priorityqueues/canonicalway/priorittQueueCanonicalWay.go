package canonicalway

import (
	"container/heap"
	"sync"
)

type PriorityQueue struct {
	sync.Mutex
	queue priorityQueueInternal
}

func (pq *PriorityQueue) Push(something interface{}, priority int) {
	pq.Lock()
	defer pq.Unlock()

	heap.Push(&pq.queue, &item{
		value:    something,
		priority: priority,
		index:    pq.queue.Len() + 1,
	})
}

func (pq *PriorityQueue) Pop() interface{} {
	pq.Lock()
	defer pq.Unlock()

	if pq.queue.Len() <= 0 {
		return nil
	}
	return heap.Pop(&pq.queue).(*item).value
}

func NewPriorityQueue() *PriorityQueue {
	pq := make(priorityQueueInternal, 0)
	heap.Init(&pq)
	return &PriorityQueue{queue: pq}
}

// An item is something we manage in a priority queue.
type item struct {
	value    interface{} // The value of the item; arbitrary.
	priority int         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A priorityQueueInternal implements heap.Interface and holds Items.
type priorityQueueInternal []*item

func (pq priorityQueueInternal) Len() int { return len(pq) }

func (pq priorityQueueInternal) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq priorityQueueInternal) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueueInternal) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueueInternal) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0: n-1]
	return item
}

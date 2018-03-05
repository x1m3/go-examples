package linkedListWay

import (
	"sync"
	"sort"
	"fmt"
)

type PriorityQueue struct {
	sync.RWMutex
	queues     map[int]*LinkedQueue
	priorities []int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{queues: make(map[int]*LinkedQueue, 16), priorities: make([]int, 0, 16)}
}

func (pq *PriorityQueue) Push(something interface{}, priority int) {
	pq.Lock()

	if q, found := pq.queues[priority]; found {
		q.Push(something)
	} else {
		fmt.Sprintf("Creando prioridad %d", priority)
		pq.queues[priority] = NewLinkedQueue()
		pq.queues[priority].Push(something)
		pq.priorities = append(pq.priorities, priority)
		sort.Sort(sort.Reverse(sort.IntSlice(pq.priorities)))
	}
	pq.Unlock()
}

func (pq *PriorityQueue) Pop() interface{} {
	pq.RLock()

	for _, priority := range pq.priorities {
		q := pq.queues[priority]
		if q.Count > 0 {
			pq.RUnlock()
			return q.Pop()
		}
	}
	pq.RUnlock()
	return nil
}

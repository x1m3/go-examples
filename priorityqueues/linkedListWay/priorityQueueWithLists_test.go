package linkedListWay

import (
	"testing"
	"fmt"
	"math/rand"
)

type Something struct {
	msg      string
	priority int
}

func TestPriorityQueue(t *testing.T) {

	const ITEMS = 1000000

	pq := NewPriorityQueue()

	for i := 0; i < ITEMS; i++ {
		priority := rand.Intn(20)
		something := &Something{msg: fmt.Sprintf("Priority %d", priority), priority: priority}
		pq.Push(something, priority)
	}

	// Take the items out; they arrive in decreasing priority order.
	expected := 1000000
	for i := 0; i < ITEMS; i++ {
		item := pq.Pop()
		if item == nil {
			t.Error("Expecting a value. Got nil")
		}
		if got := item.(*Something).priority; got >= expected {
			t.Errorf("Expecting a priority equal or lower than %d, got %d", expected, got)
		}
	}

	// Now, the queue is empty, so we expect a nil value
	if item := pq.Pop(); item != nil {
		t.Error("Expecting a nil value.")
	}
}


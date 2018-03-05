package linkedListWay

import "testing"

func TestLinkedQueueEmpty(t *testing.T) {

	l := NewLinkedQueue()

	if got, expected := l.Count, 0; got != expected {
		t.Errorf("Bad Count on empty list. Got <%d>, expected <%d>", got, expected)
	}

	empty := l.Pop()
	if empty != nil {
		t.Errorf("Bad return for empty list. Expecting nil. Got <%v>", empty)
	}
}

func TestLinkedQueueOnePushOnePop(t *testing.T) {

	l := NewLinkedQueue()
	l.Push(1)
	if got, expected := l.Count, 1; got != expected {
		t.Errorf("After inserting one element, Count is wrong. Got <%v>, expected <%v>", got, expected)
	}

	if l.first != l.last {
		t.Error("t.first shold be equal to t.last")
	}

	if got, expected := l.Pop().(int), 1; got != expected {
		t.Errorf("After push and pop one element, got <%v>, expected <%v>", got, expected)
	}

	if l.first != l.last {
		t.Error("t.first shold be equal to t.last")
	}
	if got, expected := l.Count, 0; got != expected {
		t.Errorf("Bad Count on empty list. Got <%d>, expected <%d>", got, expected)
	}

	empty := l.Pop()
	if empty != nil {
		t.Errorf("Bad return for empty list. Expecting nil. Got <%v>", empty)
	}
}

func TestLinkedQueue_PushPop(t *testing.T) {

	const ITEMS = 1000000

	l := NewLinkedQueue()

	for i := 0; i < ITEMS; i++ {
		l.Push(i)
		if got, expected := l.Count, i+1; got != expected {
			t.Errorf("Bad counter after pushing to the queue. Got <%v>, expecting <%v>", got, expected)
		}
	}

	for i := 0; i < ITEMS; i++ {
		if got, expected := l.Pop().(int), i; got != expected {
			t.Errorf("Error doing Pop from the queue. Got <%v>, expecting <%v>", got, expected)
		}
		if got, expected := l.Count, ITEMS-i-1; got != expected {
			t.Errorf("Bad counter after doing Pop from the queue. Got <%v>, expecting <%v>", got, expected)
		}
	}

	if got, expected := l.Count, 0; got != expected {
		t.Errorf("Bad Count on empty list. Got <%d>, expected <%d>", got, expected)
	}

	empty := l.Pop()
	if empty != nil {
		t.Errorf("Bad return for empty list. Expecting nil. Got <%v>", empty)
	}
}

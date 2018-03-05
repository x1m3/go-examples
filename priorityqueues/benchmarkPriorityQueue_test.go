package priorityqueues

import (
	"testing"
	"github.com/x1m3/go-examples/priorityqueues/canonicalway"
	"fmt"
	"math/rand"
	"time"
	"github.com/x1m3/go-examples/priorityqueues/linkedListWay"
)

type Something struct {
	msg      string
	priority int
}

var testData chan *Something

func init() {

	testData = make(chan *Something, 1000000)
	go func() {
		for {
			priority := rand.Intn(20)
			testData <- &Something{msg: fmt.Sprintf("Priority %d", priority), priority: priority}
		}
	}()
	time.Sleep(1 * time.Second)
}

var s1 string
var s2 string
var s3 string
var s4 string

func BenchmarkReference(b *testing.B) {

	for n := 0; n < b.N; n++ {
		for i := 0; i < 1000; i++ {
			s1 = fmt.Sprintf("lala1 %v", n)
			s2 = fmt.Sprintf("lala2 %v", n)
			s3 = fmt.Sprintf("lala3 %v", n)
			s4 = fmt.Sprintf("lala4 %v", n)
		}
	}
}

func BenchmarkPQCanonical_Push_Empty(b *testing.B) {
	benchmarkPQ_PushEmpty(canonicalway.NewPriorityQueue(), b)
}

func BenchmarkPQMyWay_Push_Empty(b *testing.B) {
	benchmarkPQ_PushEmpty(linkedListWay.NewPriorityQueue(), b)
}

func BenchmarkPQCanonical_PushPop_Empty(b *testing.B) {
	benchmarkPQ_PushPopEmpty(canonicalway.NewPriorityQueue(), b)
}

func BenchmarkPQMyWay_PushPop_Empty(b *testing.B) {
	benchmarkPQ_PushPopEmpty(linkedListWay.NewPriorityQueue(), b)
}

func benchmarkPQ_PushEmpty(pq Queue, b *testing.B) {

	for n := 0; n < b.N; n++ {
		for i := 0; i < 1000; i++ {
			item := <-testData
			pq.Push(item, item.priority)
		}
	}
}

var item *Something

func benchmarkPQ_PushPopEmpty(pq Queue, b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 1000; i++ {
			item := <-testData
			pq.Push(item, item.priority)
		}
		for i := 0; i < 1000; i++ {
			item = pq.Pop().(*Something)
		}
	}
}

package sequence

import "sync"

type Next func(interface{}) interface{}

type sequence struct {
	sync.Mutex
	last interface{}
	next Next
}

func New(fn Next, init interface{}) *sequence {
	return &sequence{last :init, next:fn}
}

func (s *sequence) Next() interface{}{
	s.Lock()
	defer s.Unlock()
	s.last = s.next
	return s.last
}

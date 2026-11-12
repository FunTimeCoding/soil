package alfa

type Queue struct {
	waiting []*waiter
}

func (q *Queue) Add() bool {
	w := &waiter{ready: make(chan struct{})}
	q.waiting = append(q.waiting, w)

	return w.granted
}

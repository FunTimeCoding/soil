package turn

import "time"

func New(bound time.Duration) *Scheduler {
	return &Scheduler{waiting: map[string][]*waiter{}, bound: bound}
}

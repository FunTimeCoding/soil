package turn

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn/waiter"
	"time"
)

func New(bound time.Duration) *Scheduler {
	return &Scheduler{waiting: map[string][]*waiter.Waiter{}, bound: bound}
}

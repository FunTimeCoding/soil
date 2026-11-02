package turn

import (
	"sync"
	"time"
)

type Scheduler struct {
	mutex   sync.Mutex
	busy    bool
	waiting map[string][]*waiter
	order   []string
	bound   time.Duration
}

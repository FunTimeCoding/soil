package turn

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn/waiter"
	"sync"
	"time"
)

type Scheduler struct {
	mutex   sync.Mutex
	busy    bool
	waiting map[string][]*waiter.Waiter
	order   []string
	bound   time.Duration
}

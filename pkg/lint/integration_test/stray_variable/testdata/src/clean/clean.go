package clean

import (
	"sync"
	"sync/atomic"
)

var (
	mutex   sync.Mutex
	once    sync.Once
	counter atomic.Int64
)

func Exercise() string {
	mutex.Lock()
	defer mutex.Unlock()
	once.Do(func() { counter.Add(1) })

	return favicon
}

package navigation

import (
	"sync"
	"time"
)

type Watcher struct {
	mutex    sync.Mutex
	started  map[string]int
	open     map[string]bool
	locators map[string]string
	changed  chan struct{}
	window   time.Duration
	limit    time.Duration
}

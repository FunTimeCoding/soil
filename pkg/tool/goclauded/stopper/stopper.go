package stopper

import "sync"

type Stopper struct {
	done chan struct{}
	once sync.Once
}

package mock_notifier

import "sync"

type Notifier struct {
	mutex       sync.Mutex
	subscribers map[chan struct{}]struct{}
	Notified    int
}

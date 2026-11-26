package navigation

import "time"

func New(
	window time.Duration,
	limit time.Duration,
) *Watcher {
	return &Watcher{
		started:  make(map[string]int),
		open:     make(map[string]bool),
		locators: make(map[string]string),
		changed:  make(chan struct{}),
		window:   window,
		limit:    limit,
	}
}

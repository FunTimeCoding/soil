package lifecycle

import "time"

func (l *Lifecycle) WithGrace(d time.Duration) *Lifecycle {
	l.grace = d

	return l
}

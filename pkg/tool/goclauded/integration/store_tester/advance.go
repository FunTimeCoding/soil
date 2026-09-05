package store_tester

import "time"

func (o *Tester) Advance(d time.Duration) {
	*o.now = o.now.Add(d)
}

package store_tester

import "time"

func (o *Tester) Clock() func() time.Time {
	return o.clock.Now
}

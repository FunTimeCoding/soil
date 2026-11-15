package store_tester

import "time"

func (o *Tester) AdvanceTo(at time.Time) {
	o.clock.AdvanceTo(at)
}

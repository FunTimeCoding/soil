package cross_service_tester

import "time"

func (o *Tester) Now() time.Time {
	return *o.now
}

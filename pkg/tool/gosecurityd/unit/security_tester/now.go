package security_tester

import "time"

func (o *Tester) now() time.Time {
	return time.Now().Add(o.Later)
}

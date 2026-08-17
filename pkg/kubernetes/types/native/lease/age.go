package lease

import "time"

func (l *Lease) Age() time.Duration {
	if !l.Held() {
		return 0
	}

	s := l.Raw.Spec

	if s.AcquireTime != nil {
		return time.Since(s.AcquireTime.Time)
	}

	if s.RenewTime != nil {
		return time.Since(s.RenewTime.Time)
	}

	return 0
}

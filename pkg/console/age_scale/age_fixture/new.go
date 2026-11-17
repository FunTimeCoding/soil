package age_fixture

import "time"

func New(v time.Duration) *Fixture {
	return &Fixture{value: v}
}

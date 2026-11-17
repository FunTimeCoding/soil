package age_fixture

import "time"

func (f *Fixture) Age() time.Duration {
	return f.value
}

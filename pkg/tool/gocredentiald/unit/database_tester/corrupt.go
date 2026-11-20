package database_tester

import "time"

func (o *Tester) Corrupt(
	path string,
	at time.Time,
) {
	o.t.Helper()
	o.replace(path, []byte("corrupt"), at)
}

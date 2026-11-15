package database_tester

import (
	"testing"
	"time"
)

func Corrupt(
	t *testing.T,
	path string,
	at time.Time,
) {
	t.Helper()
	replace(t, path, []byte("corrupt"), at)
}

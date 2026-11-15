package skip_recorder

import "testing"

func New(t testing.TB) *Recorder {
	return &Recorder{TB: t}
}

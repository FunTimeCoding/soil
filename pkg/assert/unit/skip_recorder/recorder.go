package skip_recorder

import "testing"

type Recorder struct {
	testing.TB
	skipped bool
}

package show

import "time"

type Result struct {
	ActiveState   string
	SubState      string
	ActiveEnter   time.Time
	ExecMainStart time.Time
}

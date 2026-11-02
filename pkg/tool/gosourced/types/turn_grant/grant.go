package turn_grant

import "time"

type Grant struct {
	Ahead   int
	Waited  time.Duration
	Release func()
}

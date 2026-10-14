package constant

import "time"

const ReadWriteTimeout = 10 * time.Second
const (
	TerminationGrace = 30 * time.Second
	InterruptMargin  = 5 * time.Second
)

package constant

import "time"

const (
	StreamRetryInterval = 2 * time.Second
	StreamRetryMaximum  = time.Minute
)

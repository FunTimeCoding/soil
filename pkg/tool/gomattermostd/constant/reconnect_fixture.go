package constant

import "time"

const (
	FixtureReconnectMinimum = 10 * time.Millisecond
	FixtureReconnectMaximum = 40 * time.Millisecond
)

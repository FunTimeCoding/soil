package constant

import "time"

var CruisePace = []time.Duration{
	5 * time.Minute,
	15 * time.Minute,
	30 * time.Minute,
}

const MaximumHold = 30 * time.Minute

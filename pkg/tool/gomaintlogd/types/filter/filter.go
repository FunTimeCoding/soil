package filter

import "time"

type Filter struct {
	System  string
	Service string
	User    string
	Since   time.Time
	Until   time.Time
	Limit   int
}

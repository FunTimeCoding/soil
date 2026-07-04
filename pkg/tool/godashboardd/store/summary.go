package store

import "time"

type Summary struct {
	Label string
	Count int64
	Last  time.Time
}

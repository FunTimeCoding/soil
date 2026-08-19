package store

import "time"

type Filter struct {
	Authority string
	Kind      string
	Before    *time.Time
	Revoked   *bool
	Limit     int
}

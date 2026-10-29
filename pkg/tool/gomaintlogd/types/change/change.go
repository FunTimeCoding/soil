package change

import "time"

type Change struct {
	Action      *string
	User        *string
	System      *string
	Service     *string
	Description *string
	Timestamp   *time.Time
}

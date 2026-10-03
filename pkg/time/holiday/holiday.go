package holiday

import "time"

type Holiday struct {
	month   time.Month
	day     int
	movable bool
	offset  int
}

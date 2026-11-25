package roster_entry

import "time"

type Entry struct {
	Callsign   string
	Identifier string
	LastSeen   time.Time
	OnRoster   bool
}

package roster_entry

import "time"

func New(
	callsign string,
	identifier string,
	lastSeen time.Time,
	onRoster bool,
) *Entry {
	return &Entry{
		Callsign:   callsign,
		Identifier: identifier,
		LastSeen:   lastSeen,
		OnRoster:   onRoster,
	}
}

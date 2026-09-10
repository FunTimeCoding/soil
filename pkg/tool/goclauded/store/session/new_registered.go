package session

import "time"

func NewRegistered(
	identifier string,
	callsign string,
	now time.Time,
) *Session {
	return &Session{
		Identifier: identifier,
		Name:       callsign,
		Callsign:   &callsign,
		LastSeen:   now,
		StartedAt:  now,
	}
}

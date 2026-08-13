package session

import "time"

func NewDiscovered(
	identifier string,
	now time.Time,
) *Session {
	return &Session{Identifier: identifier, LastSeen: now, StartedAt: now}
}

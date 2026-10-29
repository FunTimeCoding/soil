package event

import "time"

func New(
	t time.Time,
	process string,
	subsystem string,
	category string,
	message string,
) *Event {
	return &Event{
		Time:      t,
		Process:   process,
		Subsystem: subsystem,
		Category:  category,
		Message:   message,
	}
}

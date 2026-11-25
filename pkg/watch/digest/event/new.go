package event

import "time"

func New(
	noun string,
	author string,
	line string,
	at time.Time,
) *Event {
	return &Event{Noun: noun, Author: author, Line: line, At: at}
}

package note_writer

import "time"

func New() *Writers {
	return &Writers{
		callsign: map[int64]string{},
		written:  map[int64]time.Time{},
	}
}

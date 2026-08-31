package board_entry

import "time"

type Entry struct {
	Project           string
	ProjectIdentifier int64
	ProjectLink       string
	Reference         string
	Status            string
	Identifier        int64
	Link              string
	Updated           time.Time
}

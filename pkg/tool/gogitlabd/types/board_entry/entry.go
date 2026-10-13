package board_entry

import "time"

type Entry struct {
	Instance          string
	Project           string
	ProjectIdentifier int64
	ProjectLink       string
	Reference         string
	Status            string
	Identifier        int64
	Link              string
	Updated           time.Time
}

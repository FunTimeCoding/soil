package pipeline

import "time"

type Pipeline struct {
	Identifier        int64
	ProjectIdentifier int64
	Status            string
	Source            string
	Reference         string
	Hash              string
	Link              string
	Create            *time.Time
	Update            *time.Time
}

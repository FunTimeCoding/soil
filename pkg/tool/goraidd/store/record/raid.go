package record

import "time"

type Raid struct {
	Identifier uint
	Name       string
	Date       time.Time
	Fights     int
	Players    int
}

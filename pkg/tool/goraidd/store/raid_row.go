package store

import "time"

type RaidRow struct {
	Identifier uint
	Name       string
	Date       time.Time
	Fights     int
	Players    int
}

package top

import "time"

type Record struct {
	Name            string
	Count           int
	AverageDuration time.Duration
	CurrentlyFiring int
	Severity        string
}

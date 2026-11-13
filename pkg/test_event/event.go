package test_event

import "time"

type Event struct {
	Time        time.Time
	Action      string
	Package     string
	ImportPath  string
	Test        string
	Elapsed     float64
	Output      string
	OutputType  string
	FailedBuild string
}

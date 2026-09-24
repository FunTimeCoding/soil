package report

import "github.com/funtimecoding/soil/pkg/crap/entry"

type Report struct {
	Root    string         `json:"root"`
	Entries []*entry.Entry `json:"entries"`
	Skipped int            `json:"skipped"`
}

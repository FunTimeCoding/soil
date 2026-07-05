package reaper

import "github.com/funtimecoding/go-library/pkg/face"

func New(reporter face.Reporter) *Reaper {
	return &Reaper{
		reporter: reporter,
		stop:     make(chan struct{}),
	}
}

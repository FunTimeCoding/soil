package reaper

import "github.com/funtimecoding/soil/pkg/face"

func New(reporter face.Reporter) *Reaper {
	return &Reaper{
		reporter: reporter,
		stop:     make(chan struct{}),
	}
}

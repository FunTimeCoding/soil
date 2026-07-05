package reaper

import "github.com/funtimecoding/go-library/pkg/face"

type Reaper struct {
	reporter face.Reporter
	stop     chan struct{}
}

package reacher

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer/reacher/host"
	"time"
)

func New(r face.Reporter) *Reacher {
	return &Reacher{
		hosts:    map[string]*host.Host{},
		refused:  map[string]bool{},
		reporter: r,
		clock:    time.Now,
	}
}

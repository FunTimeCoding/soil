package outage_tracker

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer/outage_tracker/host"
	"time"
)

func New(r face.Reporter) *Tracker {
	return &Tracker{
		hosts:    map[string]*host.Host{},
		refused:  map[string]bool{},
		reporter: r,
		clock:    time.Now,
	}
}

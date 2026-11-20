package tracker_source

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/vulnerability/face"
	"time"
)

func New(
	c face.SecurityTracker,
	g *peer.Guard,
	clock func() time.Time,
	cache string,
	supported bool,
	releases ...string,
) *Source {
	return &Source{
		client:    c,
		guard:     g,
		clock:     clock,
		cache:     cache,
		supported: supported,
		releases:  releases,
	}
}

package tracker_source

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/tracker"
	"github.com/funtimecoding/soil/pkg/vulnerability/face"
	"sync/atomic"
	"time"
)

type Source struct {
	client    face.SecurityTracker
	guard     *peer.Guard
	clock     func() time.Time
	cache     string
	supported bool
	releases  []string
	tracker   atomic.Pointer[tracker.Tracker]
	verified  atomic.Pointer[time.Time]
}

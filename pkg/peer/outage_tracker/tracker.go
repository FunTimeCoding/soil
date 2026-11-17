package outage_tracker

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer/outage_tracker/host"
	"sync"
	"time"
)

type Tracker struct {
	hosts    map[string]*host.Host
	refused  map[string]bool
	reporter face.Reporter
	clock    func() time.Time
	mutex    sync.Mutex
}

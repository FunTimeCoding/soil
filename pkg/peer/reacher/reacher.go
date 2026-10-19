package reacher

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer/reacher/host"
	"sync"
	"time"
)

type Reacher struct {
	hosts    map[string]*host.Host
	refused  map[string]bool
	reporter face.Reporter
	clock    func() time.Time
	mutex    sync.Mutex
}

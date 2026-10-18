package reacher

import (
	"github.com/funtimecoding/soil/pkg/reacher/host"
	"sync"
	"time"
)

type Reacher struct {
	hosts   map[string]*host.Host
	refused map[string]bool
	clock   func() time.Time
	mutex   sync.Mutex
}

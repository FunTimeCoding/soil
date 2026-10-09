package reacher

import (
	"github.com/funtimecoding/soil/pkg/reacher/host"
	"sync"
	"time"
)

type Reacher struct {
	hosts map[string]*host.Host
	clock func() time.Time
	mutex sync.Mutex
}

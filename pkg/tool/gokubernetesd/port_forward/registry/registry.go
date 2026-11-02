package registry

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/port_forward"
	"sync"
)

type Registry struct {
	mutex    sync.Mutex
	forwards map[string]*port_forward.Forward
	running  sync.WaitGroup
}

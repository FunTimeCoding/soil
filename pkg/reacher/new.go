package reacher

import (
	"github.com/funtimecoding/soil/pkg/reacher/host"
	"time"
)

func New() *Reacher {
	return &Reacher{
		hosts:   map[string]*host.Host{},
		refused: map[string]bool{},
		clock:   time.Now,
	}
}

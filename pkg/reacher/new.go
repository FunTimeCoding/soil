package reacher

import (
	"github.com/funtimecoding/soil/pkg/reacher/host"
	"time"
)

func New() *Reacher {
	return &Reacher{hosts: map[string]*host.Host{}, clock: time.Now}
}

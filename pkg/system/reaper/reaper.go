package reaper

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/system/types/zombie_detail"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type Reaper struct {
	reporter face.Reporter
	reaped   *prometheus.CounterVec
	zombies  func() map[int]zombie_detail.Detail
	owned    sync.Map
	stop     chan struct{}
	closing  sync.Once
}

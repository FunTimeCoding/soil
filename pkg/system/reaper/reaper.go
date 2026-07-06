package reaper

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"sync"
)

type Reaper struct {
	reporter face.Reporter
	owned    sync.Map
	stop     chan struct{}
}

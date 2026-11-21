package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/debian"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/score"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/system/matched"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store"
	"sync"
	"time"
)

type Service struct {
	store             *store.Store
	debian            *debian.Service
	score             *score.Service
	clock             func() time.Time
	staleAfter        time.Duration
	trackerStaleAfter time.Duration
	cache             map[string]*matched.Matched
	mutex             sync.Mutex
}

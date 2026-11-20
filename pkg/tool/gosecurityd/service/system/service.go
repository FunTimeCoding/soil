package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/service/debian"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store"
	"time"
)

type Service struct {
	store             *store.Store
	debian            *debian.Service
	clock             func() time.Time
	staleAfter        time.Duration
	trackerStaleAfter time.Duration
}

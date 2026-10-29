package shared_reranker

import (
	"sync"
	"sync/atomic"
)

var (
	shared = sync.OnceValue(load)
	loaded atomic.Bool
)

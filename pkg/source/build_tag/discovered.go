package build_tag

import (
	"sync"
	"sync/atomic"
)

var (
	discovered sync.Map
	memoizing  atomic.Bool
)

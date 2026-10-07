package search_cache_entry

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"time"
)

type Entry struct {
	Outcome *result.Outcome
	Expiry  time.Time
}

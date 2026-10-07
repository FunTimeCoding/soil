package search_cache_entry

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"time"
)

func New(
	outcome *result.Outcome,
	expiry time.Time,
) *Entry {
	return &Entry{Outcome: outcome, Expiry: expiry}
}

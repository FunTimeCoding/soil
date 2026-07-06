package search_cache

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"time"
)

type entry struct {
	outcome *store.SearchOutcome
	expiry  time.Time
}

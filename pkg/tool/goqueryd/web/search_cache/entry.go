package search_cache

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"
	"time"
)

type entry struct {
	outcome *search.Outcome
	expiry  time.Time
}

package search

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"sync/atomic"
)

type Service struct {
	store         *store.Store
	search        *search_index.Index
	harbor        string
	searchIndexed atomic.Int64
	searchTotal   atomic.Int64
}

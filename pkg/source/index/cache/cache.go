package cache

import (
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"github.com/funtimecoding/soil/pkg/source/types/cache_entry"
	"sync"
)

type Cache struct {
	directory string
	kinds     []*kind.Kind
	lock      sync.Mutex
	entries   map[string]*cache_entry.Entry
}

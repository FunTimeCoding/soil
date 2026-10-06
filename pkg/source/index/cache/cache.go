package cache

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"sync"
)

type Cache struct {
	directory string
	kinds     []*index.Kind
	lock      sync.Mutex
	entries   map[string]*entry
}

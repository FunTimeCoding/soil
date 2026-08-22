package transcript_cache

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"sync"
)

type Cache struct {
	*claude.Client
	mutex    sync.RWMutex
	sessions map[string]*sessionEntry
	calls    map[string]*callEntry
}

package transcript_cache

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/transcript_cache/entry"
	"sync"
)

type Cache struct {
	*claude.Client
	mutex    sync.RWMutex
	sessions map[string]*entry.Session
	calls    map[string]*entry.Call
}

package session_cache

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tracker"
	"sync"
)

type Cache struct {
	entries map[string]*tracker.State
	mutex   sync.Mutex
}

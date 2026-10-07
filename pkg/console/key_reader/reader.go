package key_reader

import (
	"github.com/funtimecoding/soil/pkg/console/types/key_callback"
	"sync"
	"time"
)

type Reader struct {
	handlers map[rune]key_callback.Callback
	pressed  map[rune]time.Time
	mutex    sync.RWMutex
}

package key_reader

import (
	"sync"
	"time"
)

type Reader struct {
	handlers map[rune]Callback
	pressed  map[rune]time.Time
	mutex    sync.RWMutex
}

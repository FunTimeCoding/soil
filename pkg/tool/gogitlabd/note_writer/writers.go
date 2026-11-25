package note_writer

import (
	"sync"
	"time"
)

type Writers struct {
	mutex    sync.Mutex
	callsign map[int64]string
	written  map[int64]time.Time
}

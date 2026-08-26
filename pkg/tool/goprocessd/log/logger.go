package log

import (
	"sync"
	"time"
)

type Logger struct {
	colorIndex   int
	name         string
	writes       chan []byte
	done         chan struct{}
	timeout      time.Duration
	buffers      [][]byte
	maxNameWidth    int
	history         []string
	generationStart int
}

var mutex sync.Mutex

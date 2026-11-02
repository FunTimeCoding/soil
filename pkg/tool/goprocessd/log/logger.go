package log

import (
	"io"
	"sync"
	"time"
)

type Logger struct {
	output          io.Writer
	colorIndex      int
	name            string
	timeout         time.Duration
	buffers         [][]byte
	bufferMutex     sync.Mutex
	flush           *time.Timer
	maxNameWidth    int
	history         []string
	generationStart int
}

var mutex sync.Mutex

package log

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"time"
)

func New(
	name string,
	colorIndex int,
	maxNameWidth int,
) *Logger {
	l := &Logger{
		colorIndex:   colorIndex % len(constant.Colors),
		name:         name,
		writes:       make(chan []byte),
		done:         make(chan struct{}),
		timeout:      2 * time.Millisecond,
		maxNameWidth: maxNameWidth,
	}
	go l.writeLines()

	return l
}

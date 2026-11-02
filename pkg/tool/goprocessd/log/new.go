package log

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"io"
	"time"
)

func New(
	output io.Writer,
	name string,
	colorIndex int,
	maxNameWidth int,
) *Logger {
	return &Logger{
		output:       output,
		colorIndex:   colorIndex % len(constant.Colors),
		name:         name,
		timeout:      2 * time.Millisecond,
		maxNameWidth: maxNameWidth,
	}
}

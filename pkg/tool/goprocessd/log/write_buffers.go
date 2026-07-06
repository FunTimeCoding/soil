package log

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"os"
	"strings"
	"time"
)

func (l *Logger) writeBuffers(line []byte) {
	mutex.Lock()
	now := time.Now().Format("15:04:05")
	writer.Print(
		os.Stdout,
		"\x1b[%dm%s %*s | \x1b[m",
		colors[l.colorIndex],
		now,
		l.maxNameWidth,
		l.name,
	)
	l.buffers = append(l.buffers, line)

	for _, b := range l.buffers {
		_, e := os.Stdout.Write(b)
		errors.PanicOnError(e)
	}

	l.buffers = l.buffers[:0]
	text := strings.TrimRight(string(line), "\n")

	if len(text) > 0 {
		if len(l.history) >= constant.HistoryCapacity {
			l.history = l.history[1:]
		}

		l.history = append(l.history, text)
	}

	mutex.Unlock()
}

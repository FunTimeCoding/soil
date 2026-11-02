package log

import "time"

func (l *Logger) scheduleFlush() {
	l.stopFlush()
	l.flush = time.AfterFunc(l.timeout, l.flushPending)
}

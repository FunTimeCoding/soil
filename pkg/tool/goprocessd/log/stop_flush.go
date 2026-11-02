package log

func (l *Logger) stopFlush() {
	if l.flush != nil {
		l.flush.Stop()
		l.flush = nil
	}
}

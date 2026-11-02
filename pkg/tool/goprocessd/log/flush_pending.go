package log

func (l *Logger) flushPending() {
	l.bufferMutex.Lock()
	defer l.bufferMutex.Unlock()

	if len(l.buffers) > 0 {
		l.writeBuffers([]byte("\n"))
	}

	l.flush = nil
}

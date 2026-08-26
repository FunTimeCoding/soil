package log

func (l *Logger) Clear() {
	mutex.Lock()
	defer mutex.Unlock()
	l.history = nil
	l.generationStart = 0
}

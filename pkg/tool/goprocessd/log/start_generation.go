package log

func (l *Logger) StartGeneration() {
	mutex.Lock()
	defer mutex.Unlock()
	l.generationStart = len(l.history)
}
